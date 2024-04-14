package models

import "time"

type Invitation struct {
	ID         string    `json:"id" firestore:"-"`
	FromUserID string    `json:"from_user_id" firestore:"from_user_id"`
	ToUserID   string    `json:"to_user_id" firestore:"to_user_id"`
	Status     string    `json:"status" firestore:"status"`
	SentDate   time.Time `json:"sent_date" firestore:"sent_date"`
}

const InvitationPending = "pending"