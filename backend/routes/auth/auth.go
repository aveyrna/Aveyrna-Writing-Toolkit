// backend/routes/auth/auth.go
package auth

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"backend/db"
	"backend/models"

	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
)

type creds struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type authResponse struct {
	ID       string `json:"id"` // = PublicID (uuid) exposé via json:"id" dans models.User
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/register", registerUser)
	r.Post("/login", loginUser)
	r.Get("/me", me)
	r.Post("/logout", logoutUser)
	return r
}

// ---- Token HMAC très simple (dev)
func secret() []byte {
	s := os.Getenv("AUTH_SECRET")
	if s == "" {
		s = "dev-secret-change-me"
	}
	return []byte(s)
}

// on encode: publicUUID|email|unixTs + signature HMAC, puis base64url
func makeToken(u models.User) (string, error) {
	payload := fmt.Sprintf("%s|%s|%d", u.PublicID.String(), u.Email, time.Now().Unix())
	h := hmac.New(sha256.New, secret())
	h.Write([]byte(payload))
	sig := h.Sum(nil)
	raw := payload + "|" + base64.RawURLEncoding.EncodeToString(sig)
	return base64.RawURLEncoding.EncodeToString([]byte(raw)), nil
}
func parseToken(tok string) (publicID string, email string, err error) {
	dec, err := base64.RawURLEncoding.DecodeString(tok)
	if err != nil {
		return "", "", errors.New("invalid token")
	}
	parts := strings.Split(string(dec), "|")
	if len(parts) != 4 {
		return "", "", errors.New("invalid token")
	}
	payload := strings.Join(parts[:3], "|")
	sigGiven, err := base64.RawURLEncoding.DecodeString(parts[3])
	if err != nil {
		return "", "", errors.New("invalid token")
	}
	h := hmac.New(sha256.New, secret())
	h.Write([]byte(payload))
	if !hmac.Equal(h.Sum(nil), sigGiven) {
		return "", "", errors.New("invalid token signature")
	}
	// payload: publicUUID|email|ts
	var ts int64
	if _, err := fmt.Sscanf(payload, "%s|%s|%d", &publicID, &email, &ts); err != nil {
		return "", "", errors.New("invalid token payload")
	}
	return publicID, email, nil
}

func logoutUser(w http.ResponseWriter, r *http.Request) {
	// Si un cookie HttpOnly est utilisé dans une autre conf, on le “supprime” proprement.
	http.SetCookie(w, &http.Cookie{
		Name:     "auth",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})
	w.WriteHeader(http.StatusNoContent) // 204
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var body creds
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}
	if body.Email == "" || body.Password == "" {
		http.Error(w, "email and password required", http.StatusBadRequest)
		return
	}

	println("Trying to register:  ", body.Username)

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "hash error", http.StatusInternalServerError)
		return
	}

	var u models.User
	// On retourne les champs publics (id = public_id via json tags)
	err = db.Pool.QueryRow(ctx, `
		INSERT INTO users (username, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id, public_id, username, email, created_at, updated_at`,
		body.Username, body.Email, string(hash),
	).Scan(&u.ID, &u.PublicID, &u.Username, &u.Email, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		http.Error(w, "email exists or db error", http.StatusBadRequest)
		return
	}

	tok, _ := makeToken(u)
	writeJSON(w, authResponse{
		ID:       u.PublicID.String(),
		Username: u.Username,
		Email:    u.Email,
		Token:    tok,
	})
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var body creds
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}
	if body.Email == "" || body.Password == "" {
		http.Error(w, "email and password required", http.StatusBadRequest)
		return
	}

	var u models.User
	var passwordHash string
	err := db.Pool.QueryRow(ctx, `
		SELECT id, public_id, username, email, password_hash, created_at, updated_at
		FROM users WHERE email = $1`, body.Email,
	).Scan(&u.ID, &u.PublicID, &u.Username, &u.Email, &passwordHash, &u.CreatedAt, &u.UpdatedAt)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(body.Password)) != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	tok, _ := makeToken(u)
	writeJSON(w, authResponse{
		ID:       u.PublicID.String(),
		Username: u.Username,
		Email:    u.Email,
		Token:    tok,
	})
}

func me(w http.ResponseWriter, r *http.Request) {
	authz := r.Header.Get("Authorization")
	if authz == "" || !strings.HasPrefix(authz, "Bearer ") {
		http.Error(w, "missing bearer", http.StatusUnauthorized)
		return
	}
	publicID, email, err := parseToken(strings.TrimPrefix(authz, "Bearer "))
	if err != nil {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}

	var u models.User
	if err := db.Pool.QueryRow(context.Background(), `
		SELECT id, public_id, username, email, created_at, updated_at
		FROM users WHERE public_id = $1 AND email = $2`,
		publicID, email,
	).Scan(&u.ID, &u.PublicID, &u.Username, &u.Email, &u.CreatedAt, &u.UpdatedAt); err != nil {
		http.Error(w, "user not found", http.StatusUnauthorized)
		return
	}

	// On renvoie directement ton modèle: ID interne et Password sont masqués par les tags json
	writeJSON(w, u)
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}
