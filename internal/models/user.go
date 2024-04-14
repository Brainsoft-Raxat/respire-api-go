package models

import "time"

type User struct {
	ID    string `json:"id" firestore:"-"`
	Name  string `json:"name" firestore:"name"`
	Email string `json:"email" firestore:"email"`
	// PasswordHash   string           `json:"-"`

	QuitDate       time.Time        `json:"quite_date" firestore:"quite_date"`
	CurrentStreak  int              `json:"current_streak" firestore:"current_streak"`
	LongestStreak  int              `json:"longest_streak" firestore:"longest_streak"`
	// SmokingHistory []SmokingSession `json:"smoking_history" firestore:"smoking_history"`

	//Friends     int `json:"friends" firestore:"friends"`
	//Invitations int `json:"invitations" firestore:"invitations"`

	CreatedAt time.Time `json:"created_at" firestore:"created_at"`
	UpdatedAt time.Time `json:"updated_at" firestore:"updated_at"`
}

// type SmokingSession struct {
// 	Date     time.Time `json:"date" firestore:"date"`
// 	Quantity int       `json:"quantity" firestore:"quantity"`
// }
