// backend/routes/auth/auth.go
package auth

import (
	"context"
	"crypto/hmac"
	"crypto/rand"
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

// ---- Token HMAC simple (legacy, pour compatibilité)
func secret() []byte {
	s := os.Getenv("AUTH_SECRET")
	if s == "" {
		s = "dev-secret-change-me"
	}
	return []byte(s)
}

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
	var ts int64
	if _, err := fmt.Sscanf(payload, "%s|%s|%d", &publicID, &email, &ts); err != nil {
		return "", "", errors.New("invalid token payload")
	}
	return publicID, email, nil
}

// ---- Helpers sessions opaques
func randomToken(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}
func sha256b64(s string) string {
	sum := sha256.Sum256([]byte(s))
	return base64.RawURLEncoding.EncodeToString(sum[:])
}

func logoutUser(w http.ResponseWriter, r *http.Request) {
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
	w.WriteHeader(http.StatusNoContent)
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

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "hash error", http.StatusInternalServerError)
		return
	}

	var u models.User
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

	// 1) Crée un token de session
	sessionTok, err := randomToken(32)
	if err != nil {
		http.Error(w, "token error", 500)
		return
	}
	tokHash := sha256b64(sessionTok)

	var expiresAt time.Time
	err = db.Pool.QueryRow(ctx, `
		INSERT INTO sessions (user_id, token_hash, expires_at)
		VALUES ($1, $2, now() + interval '30 days')
		RETURNING expires_at`, u.ID, tokHash).Scan(&expiresAt)
	if err != nil {
		http.Error(w, "session error", 500)
		return
	}

	// 2) Pose le cookie
	const sessionDays = 30
	dur := time.Hour * 24 * sessionDays
	exp := time.Now().Add(dur)

	http.SetCookie(w, &http.Cookie{
		Name:     "auth",
		Value:    sessionTok,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,                // en dev: pas de Secure
		SameSite: http.SameSiteLaxMode, // http.SameSiteNoneMode,
		Expires:  exp,
		MaxAge:   int(dur.Seconds()),
	})

	// 3) Réponse JSON (compatibilité avec ton front actuel)
	tok, _ := makeToken(u)
	writeJSON(w, authResponse{
		ID:       u.PublicID.String(),
		Username: u.Username,
		Email:    u.Email,
		Token:    tok,
	})
}

func me(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// 1) Tentative via cookie de session opaque
	if c, err := r.Cookie("auth"); err == nil && c.Value != "" {
		tokHash := sha256b64(c.Value)

		var u models.User
		err := db.Pool.QueryRow(ctx, `
			SELECT u.id, u.public_id, u.username, u.email, u.created_at, u.updated_at
			FROM sessions s
			JOIN users u ON u.id = s.user_id
			WHERE s.token_hash = $1 AND s.expires_at > now()
			LIMIT 1
		`, tokHash).Scan(&u.ID, &u.PublicID, &u.Username, &u.Email, &u.CreatedAt, &u.UpdatedAt)

		if err == nil {
			writeJSON(w, u)
			return
		}
		// sinon on tente le fallback legacy
	}

	// 2) Fallback legacy: Authorization: Bearer <token HMAC>
	authz := r.Header.Get("Authorization")
	if strings.HasPrefix(authz, "Bearer ") {
		token := strings.TrimPrefix(authz, "Bearer ")
		publicID, email, err := parseToken(token)
		if err == nil {
			var u models.User
			if err := db.Pool.QueryRow(ctx, `
				SELECT id, public_id, username, email, created_at, updated_at
				FROM users
				WHERE public_id = $1 AND email = $2
			`, publicID, email).Scan(&u.ID, &u.PublicID, &u.Username, &u.Email, &u.CreatedAt, &u.UpdatedAt); err == nil {
				writeJSON(w, u)
				return
			}
		}
	}

	http.Error(w, "unauthorized", http.StatusUnauthorized)
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}
