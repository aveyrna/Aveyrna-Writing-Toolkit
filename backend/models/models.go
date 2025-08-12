package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        int       `json:"-"`
	PublicID  uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // à ne jamais exposer dans une API
	CreatedAt time.Time `json:"created_at"`
}

type Project struct {
	ID           int       `json:"-"`
	PublicID     uuid.UUID `json:"id"`
	UserID       int       `json:"user_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	StoryModelID *int      `json:"story_model_id,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}

type Character struct {
	ID               int       `json:"-"`
	PublicID         uuid.UUID `json:"id"`
	ProjectID        int       `json:"project_id"`
	Name             string    `json:"name"`
	Role             string    `json:"role"`
	Bio              string    `json:"bio"`
	Background       string    `json:"background"`
	Personality      string    `json:"personality"`
	Objective        string    `json:"objective"`
	InternalConflict string    `json:"internal_conflict"`
	ArcType          string    `json:"arc_type"`
	Notes            string    `json:"notes"`
	AvatarURL        string    `json:"avatar_url"`
}

type Location struct {
	ID           int       `json:"-"`
	PublicID     uuid.UUID `json:"id"`
	ProjectID    int       `json:"project_id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	MapReference string    `json:"map_reference"`
	ImageURL     string    `json:"image_url"`
}

type Chapter struct {
	ID           int       `json:"-"`
	PublicID     uuid.UUID `json:"id"`
	ProjectID    int       `json:"project_id"`
	Title        string    `json:"title"`
	Synopsis     string    `json:"synopsis"`
	StoryPhaseID *int      `json:"story_phase_id,omitempty"`
	OrderIndex   int       `json:"order_index"`
}

type Scene struct {
	ID              int       `json:"-"`
	PublicID        uuid.UUID `json:"id"`
	ChapterUUID     uuid.UUID `json:"chapter_uuid"`
	ChapterPublicID string    `json:"chapter_public_id"`
	Title           string    `json:"title"`
	Content         string    `json:"content"`
	Summary         string    `json:"summary"`
	LocationID      *int      `json:"location_id,omitempty"`
	OrderIndex      int       `json:"order_index"`
}

type Faction struct {
	ID          int       `json:"-"`
	PublicID    uuid.UUID `json:"id"`
	ProjectID   int       `json:"project_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Color       string    `json:"color"`
}

type FullProject struct {
	Project    Project     `json:"project"`
	Characters []Character `json:"characters"`
	Locations  []Location  `json:"locations"`
	Chapters   []Chapter   `json:"chapters"`
	Scenes     []Scene     `json:"scenes"`
	Factions   []Faction   `json:"factions"`
}
