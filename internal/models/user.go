package models

import "time"

type User struct {
	ID       string `json:"id" firestore:"id,omitempty"`
	Name     string `json:"name" firestore:"name,omitempty"`
	Email    string `json:"email" firestore:"email,omitempty"`
	Username string `json:"username" firestore:"username,omitempty"`
	Avatar   string `json:"avatar" firestore:"avatar,omitempty"`
	// PasswordHash   string           `json:"-"`

	QuitDate      time.Time `json:"quite_date" firestore:"quite_date,omitempty"`
	LongestStreak int       `json:"longest_streak" firestore:"longest_streak,omitempty"`
	Status        string    `json:"status" firestore:"status,omitempty"`

	CreatedAt time.Time `json:"created_at" firestore:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at" firestore:"updated_at,omitempty"`
}

type ShortUser struct {
	ID         string `json:"id" firestore:"-"`
	Name       string `json:"name" firestore:"name,omitempty"`
	Username   string `json:"username" firestore:"username,omitempty"`
	Avatar     string `json:"avatar" firestore:"avatar,omitempty"`
	SmokeCount int    `json:"smoke_count,omitempty" firestore:"smoke_count,omitempty"`
	Position   int    `json:"position,omitempty" firestore:"position,omitempty"`
}

const (
	STATUS_CREATED   = "CREATED"
	STATUS_COMPLETED = "COMPLETED"
)

// type SmokingSession struct {
// 	Date     time.Time `json:"date" firestore:"date"`
// 	Quantity int       `json:"quantity" firestore:"quantity"`
// }
