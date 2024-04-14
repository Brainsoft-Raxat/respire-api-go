package models

import "time"

type Challenge struct {
	ID          string    `json:"id" firestore:"-"`
	Name        string    `json:"name" firestore:"name"`
	StartDate   time.Time `json:"start_date" firestore:"start_date"`
	EndDate     time.Time `json:"end_date" firestore:"end_date"`
	Description string    `json:"description" firestore:"description"`
	Tasks       []Task    `json:"tasks" firestore:"tasks"`
	Prize       string    `json:"prize" firestore:"prize"`
	Penalty     int       `json:"penalty" firestore:"penalty"`
	IsActive    bool      `json:"is_active" firestore:"is_active"`
}

