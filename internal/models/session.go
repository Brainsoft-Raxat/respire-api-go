package models

import "time"

type SmokeSession struct {
	ID        string    `json:"id" firestore:"id"`
	UID       string    `json:"user_id" firestore:"user_id"`
	Timestamp time.Time `json:"timestamp" firestore:"timestamp"`
	Count     int       `json:"count" firestore:"count"`
}
