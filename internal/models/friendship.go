package models

import "time"

type Friendship struct {
	ID    string    `json:"id" firestore:"-"`
	Users []string  `json:"users" firestore:"users"`
	Since time.Time `json:"since" firestore:"since"`
}

type Friend struct {
	ID    string    `json:"id"`
	Name  string    `json:"name"`
	Since time.Time `json:"since"`
}