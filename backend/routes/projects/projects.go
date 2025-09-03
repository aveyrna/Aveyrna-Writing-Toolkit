package projects

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"backend/db"
	"backend/models"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", getAllProjects)
	r.Get("/user/{userID}", getProjectsByUser)
	r.Post("/", createProject)
	r.Get("/{id}", getProjectByID)
	r.Put("/{id}", updateProject)
	r.Delete("/{id}", deleteProject)
	r.Get("/{id}/full", getFullProject)
	r.Get("/public/{uuid}/full", getFullProjectByUUID)
	r.Get("/user/{userID}/full", getFullProjectsByUser)

	return r
}

func getAllProjects(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Pool.Query(context.Background(),
		`SELECT id, user_id, title, description, story_model_id, created_at FROM projects`)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var p models.Project
		err := rows.Scan(&p.ID, &p.UserID, &p.Title, &p.Description, &p.StoryModelID, &p.CreatedAt)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		projects = append(projects, p)
	}

	json.NewEncoder(w).Encode(projects)
}

func getProjectsByUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	rows, err := db.Pool.Query(context.Background(),
		`SELECT id, user_id, title, description, story_model_id, created_at
		 FROM projects WHERE user_id = $1`, userID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var p models.Project
		err := rows.Scan(&p.ID, &p.UserID, &p.Title, &p.Description, &p.StoryModelID, &p.CreatedAt)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		projects = append(projects, p)
	}

	json.NewEncoder(w).Encode(projects)
}

func getProjectByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var p models.Project
	err := db.Pool.QueryRow(context.Background(),
		`SELECT id, user_id, title, description, story_model_id, created_at
		 FROM projects WHERE id = $1`, id).
		Scan(&p.ID, &p.UserID, &p.Title, &p.Description, &p.StoryModelID, &p.CreatedAt)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	json.NewEncoder(w).Encode(p)
}

// func createProject(w http.ResponseWriter, r *http.Request) {
// 	var p models.Project
// 	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
// 		http.Error(w, err.Error(), 400)
// 		return
// 	}

// 	err := db.Pool.QueryRow(context.Background(),
// 		`INSERT INTO projects (user_id, title, description, story_model_id)
// 		 VALUES ($1, $2, $3, $4) RETURNING id, created_at`,
// 		p.UserID, p.Title, p.Description, p.StoryModelID).
// 		Scan(&p.ID, &p.CreatedAt)
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(p)
// }

func createProject(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// 1) Récupère l'utilisateur depuis la session (cookie "auth")
	c, err := r.Cookie("auth")
	if err != nil || c.Value == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	// hash du token comme dans /auth
	sum := sha256.Sum256([]byte(c.Value))
	tokHash := base64.RawURLEncoding.EncodeToString(sum[:])

	var userID int64
	if err := db.Pool.QueryRow(ctx, `
		SELECT u.id
		FROM sessions s
		JOIN users u ON u.id = s.user_id
		WHERE s.token_hash = $1 AND s.expires_at > now()
		LIMIT 1
	`, tokHash).Scan(&userID); err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	// 2) Decode payload minimal (NE PREND PAS user_id du client)
	var body struct {
		Title        string `json:"title"`
		Description  string `json:"description"`
		StoryModelID *int64 `json:"story_model_id,omitempty"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(body.Title) == "" {
		http.Error(w, "title requis", http.StatusBadRequest)
		return
	}

	// 3) Insert côté DB en générant public_id
	var p struct {
		ID           int64     `json:"id"`
		PublicID     uuid.UUID `json:"public_id"`
		UserID       int64     `json:"user_id"`
		Title        string    `json:"title"`
		Description  string    `json:"description"`
		StoryModelID *int64    `json:"story_model_id,omitempty"`
		CreatedAt    time.Time `json:"created_at"`
	}
	err = db.Pool.QueryRow(ctx, `
		INSERT INTO projects (public_id, user_id, title, description, story_model_id, created_at)
		VALUES (gen_random_uuid(), $1, $2, COALESCE($3,''), $4, now())
		RETURNING id, public_id, user_id, title, description, story_model_id, created_at
	`, userID, body.Title, body.Description, body.StoryModelID).
		Scan(&p.ID, &p.PublicID, &p.UserID, &p.Title, &p.Description, &p.StoryModelID, &p.CreatedAt)
	if err != nil {
		http.Error(w, "db error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 4) Réponse
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(p)
}

func updateProject(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var p models.Project
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	_, err := db.Pool.Exec(context.Background(),
		`UPDATE projects SET title = $1, description = $2, story_model_id = $3 WHERE id = $4`,
		p.Title, p.Description, p.StoryModelID, id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func deleteProject(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	_, err := db.Pool.Exec(context.Background(),
		`DELETE FROM projects WHERE id = $1`, id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func getFullProject(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	projectID := chi.URLParam(r, "id")
	fmt.Println("🔍 getFullProject ID =", projectID)

	var full models.FullProject

	// Chargement du projet
	err := db.Pool.QueryRow(ctx,
		`SELECT id, public_id, user_id, title, description, story_model_id, created_at
		FROM projects WHERE id = $1`, projectID).
		Scan(&full.Project.ID, &full.Project.PublicID, &full.Project.UserID,
			&full.Project.Title, &full.Project.Description,
			&full.Project.StoryModelID, &full.Project.CreatedAt)
	if err != nil {
		http.Error(w, "Project not found", 404)
		fmt.Println("❌ project query failed:", err)
		return
	}
	fmt.Println("✅ Project loaded:", full.Project.Title)

	// Characters
	full.Characters, err = getCharactersByProjectID(ctx, projectID)
	if err != nil {
		http.Error(w, "Error loading characters", 500)
		fmt.Println("❌ getCharactersByProjectID error:", err)
		return
	}
	fmt.Println("✅ Characters loaded:", len(full.Characters))

	// Locations
	full.Locations, err = getLocationsByProjectID(ctx, projectID)
	if err != nil {
		http.Error(w, "Error loading locations", 500)
		fmt.Println("❌ getLocationsByProjectID error:", err)
		return
	}
	fmt.Println("✅ Locations loaded:", len(full.Locations))

	// Chapters
	full.Chapters, err = getChaptersByProjectID(ctx, projectID)
	if err != nil {
		http.Error(w, "Error loading chapters", 500)
		fmt.Println("❌ getChaptersByProjectID error:", err)
		return
	}
	fmt.Println("✅ Chapters loaded:", len(full.Chapters))

	// Scenes
	full.Scenes, err = getScenesByProjectID(ctx, projectID)
	if err != nil {
		http.Error(w, "Error loading scenes", 500)
		fmt.Println("❌ getScenesByProjectID error:", err)
		return
	}
	fmt.Println("✅ Scenes loaded:", len(full.Scenes))

	// Factions
	full.Factions, err = getFactionsByProjectID(ctx, projectID)
	if err != nil {
		http.Error(w, "Error loading factions", 500)
		fmt.Println("❌ getFactionsByProjectID error:", err)
		return
	}
	fmt.Println("✅ Factions loaded:", len(full.Factions))

	// Encode JSON
	err = json.NewEncoder(w).Encode(full)
	if err != nil {
		http.Error(w, "JSON encoding failed", 500)
		fmt.Println("❌ JSON encode error:", err)
	}
}

func getCharactersByProjectID(ctx context.Context, projectID string) ([]models.Character, error) {
	rows, err := db.Pool.Query(ctx, `
		SELECT id, public_id, project_id, name, role, bio, background, personality,
		       objective, internal_conflict, arc_type, notes, avatar_url
		FROM characters WHERE project_id = $1`, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var characters []models.Character
	for rows.Next() {
		var c models.Character
		if err := rows.Scan(&c.ID, &c.PublicID, &c.ProjectID, &c.Name, &c.Role, &c.Bio,
			&c.Background, &c.Personality, &c.Objective, &c.InternalConflict,
			&c.ArcType, &c.Notes, &c.AvatarURL); err != nil {
			return nil, err
		}
		characters = append(characters, c)
	}
	return characters, nil
}

func getLocationsByProjectID(ctx context.Context, projectID string) ([]models.Location, error) {
	rows, err := db.Pool.Query(ctx, `
		SELECT id, public_id, project_id, name, description, map_reference, image_url
		FROM locations WHERE project_id = $1`, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Location
	for rows.Next() {
		var l models.Location
		if err := rows.Scan(&l.ID, &l.PublicID, &l.ProjectID, &l.Name,
			&l.Description, &l.MapReference, &l.ImageURL); err != nil {
			return nil, err
		}
		list = append(list, l)
	}
	return list, nil
}

func getChaptersByProjectID(ctx context.Context, projectID string) ([]models.Chapter, error) {
	rows, err := db.Pool.Query(ctx, `
		SELECT id, public_id, project_id, title, synopsis, story_phase_id, order_index
		FROM chapters WHERE project_id = $1 ORDER BY order_index ASC`, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Chapter
	for rows.Next() {
		var c models.Chapter
		if err := rows.Scan(&c.ID, &c.PublicID, &c.ProjectID, &c.Title,
			&c.Synopsis, &c.StoryPhaseID, &c.OrderIndex); err != nil {
			return nil, err
		}
		list = append(list, c)
	}
	return list, nil
}

func getScenesByProjectID(ctx context.Context, projectID string) ([]models.Scene, error) {
	rows, err := db.Pool.Query(ctx, `
		SELECT s.id, s.public_id, s.chapter_uuid, s.title, s.content, s.summary, s.location_id, s.order_index
		FROM scenes s
		INNER JOIN chapters c ON s.chapter_id = c.id
		WHERE c.project_id = $1 ORDER BY s.order_index ASC`, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Scene
	for rows.Next() {
		var s models.Scene
		if err := rows.Scan(&s.ID, &s.PublicID, &s.ChapterUUID, &s.Title,
			&s.Content, &s.Summary, &s.LocationID, &s.OrderIndex); err != nil {
			return nil, err
		}
		list = append(list, s)
	}
	return list, nil
}

func getFactionsByProjectID(ctx context.Context, projectID string) ([]models.Faction, error) {
	rows, err := db.Pool.Query(ctx, `
		SELECT id, public_id, project_id, name, description, color
		FROM factions WHERE project_id = $1`, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Faction
	for rows.Next() {
		var f models.Faction
		if err := rows.Scan(&f.ID, &f.PublicID, &f.ProjectID, &f.Name,
			&f.Description, &f.Color); err != nil {
			return nil, err
		}
		list = append(list, f)
	}
	return list, nil
}

func getFullProjectByUUID(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	uuidStr := chi.URLParam(r, "uuid")
	fmt.Println("▶️ UUID reçu :", uuidStr)

	publicID, err := uuid.Parse(uuidStr)
	if err != nil {
		http.Error(w, "Invalid UUID", 400)
		fmt.Println("❌ UUID parsing error:", err)
		return
	}

	var projectID int
	err = db.Pool.QueryRow(ctx,
		`SELECT id FROM projects WHERE public_id = $1`, publicID).Scan(&projectID)
	if err != nil {
		http.Error(w, "Project not found", 404)
		fmt.Println("❌ DB lookup error:", err)
		return
	}
	fmt.Println("✅ Projet trouvé, ID =", projectID)

	// Appel du handler "normal"
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", strconv.Itoa(projectID))
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

	getFullProject(w, r)
}

func getFullProjectsByUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	userID := chi.URLParam(r, "userID")
	fmt.Println("🔍 Chargement projets complets pour userID:", userID)

	var userDbId int
	err := db.Pool.QueryRow(ctx,
		`SELECT id FROM users WHERE public_id = $1`, userID).Scan(&userDbId)
	if err != nil {
		http.Error(w, "User not found", 404)
		fmt.Println("❌ User not found:", err)
		return
	}

	fmt.Println("✅ User found, DB ID =", userDbId)

	rows, err := db.Pool.Query(ctx,
		`SELECT id, public_id, user_id, title, description, story_model_id, created_at
		 FROM projects WHERE user_id = $1`, userDbId)
	if err != nil {
		http.Error(w, "DB error", 500)
		fmt.Println("❌ DB error:", err)
		return
	}
	defer rows.Close()

	var fullProjects []models.FullProject

	for rows.Next() {
		var p models.Project
		err := rows.Scan(&p.ID, &p.PublicID, &p.UserID, &p.Title, &p.Description, &p.StoryModelID, &p.CreatedAt)
		if err != nil {
			http.Error(w, "Scan error", 500)
			fmt.Println("❌ Scan project error:", err)
			return
		}

		full := models.FullProject{Project: p}

		full.Characters, err = getCharactersByProjectID(ctx, fmt.Sprint(p.ID))
		if err != nil {
			http.Error(w, "Characters error", 500)
			fmt.Println("❌ getCharacters:", err)
			return
		}

		full.Locations, err = getLocationsByProjectID(ctx, fmt.Sprint(p.ID))
		if err != nil {
			http.Error(w, "Locations error", 500)
			fmt.Println("❌ getLocations:", err)
			return
		}

		full.Chapters, err = getChaptersByProjectID(ctx, fmt.Sprint(p.ID))
		if err != nil {
			http.Error(w, "Chapters error", 500)
			fmt.Println("❌ getChapters:", err)
			return
		}

		full.Scenes, err = getScenesByProjectID(ctx, fmt.Sprint(p.ID))
		if err != nil {
			http.Error(w, "Scenes error", 500)
			fmt.Println("❌ getScenes:", err)
			return
		}

		full.Factions, err = getFactionsByProjectID(ctx, fmt.Sprint(p.ID))
		if err != nil {
			http.Error(w, "Factions error", 500)
			fmt.Println("❌ getFactions:", err)
			return
		}

		fullProjects = append(fullProjects, full)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Println("✅ Résultat retourné :", fullProjects)

	json.NewEncoder(w).Encode(fullProjects)
}
