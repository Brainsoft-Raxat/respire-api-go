package models

import "time"

type Friendship struct {
	ID       string    `json:"id" firestore:"-"`
	UserID   string    `json:"user_id" firestore:"user_id"`
	FriendID string    `json:"friend_id" firestore:"friend_id"`
	Since    time.Time `json:"since" firestore:"since"`
}
