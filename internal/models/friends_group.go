package models

import "time"

type FriendsGroup struct {
	ID         string      `json:"id" firestore:"-"`
	Name       string      `json:"name" firestore:"name"`
	OwnerID    string      `json:"owner_id" firestore:"owner_id"`
	Friends    []string    `json:"friends" firestore:"friends"`
	Challenges []Challenge `json:"challenges" firestore:"challenges"`
}

type FriendsGroupInvitation struct {
	ID             string    `json:"id" firestore:"-"`
	FriendsGroupID string    `json:"friends_group_id" firestore:"friends_group_id"`
	FromUserID     string    `json:"from_user_id" firestore:"from_user_id"`
	ToUserID       string    `json:"to_user_id" firestore:"to_user_id"`
	Status         string    `json:"status" firestore:"status"`
	SentDate       time.Time `json:"sent_date" firestore:"sent_date"`
}

type CreateFriendsGroupRequest struct {
	Name    string   `json:"name"`
	Invites []string `json:"invites"`
}

type UpdateFriendsGroupRequest struct {
    Name string `json:"name"`
}