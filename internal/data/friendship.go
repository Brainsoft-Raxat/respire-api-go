package data

import "github.com/Brainsoft-Raxat/respire-api-go/internal/models"

type CreateFriendshipInvitationRequest struct {
	FriendID string `json:"friend_id"`
}

type CreateFriendshipInvitationResponse Accepted

type GetFriendshipInvitationsRequest struct {
	FilterRequest
}

type GetFriendshipInvitationsResponse struct {
	Invitations []*models.Invitation `json:"invitations"`
}

type HandleFriendshipInvitationRequest struct {
	InvitationID string `json:"invitation_id"`
	Accept       bool   `json:"accept"`
}

type HandleFriendshipInvitationResponse Accepted

type GetFriendsListRequest struct {
	UserID string `json:"user_id"`
}

type GetFriendsListResponse struct {
	Friends []*models.ShortUser `json:"friends"`
}

type RemoveFriendRequest struct {
	FriendID string `json:"friend_id"`
}

type RemoveFriendResponse Accepted
