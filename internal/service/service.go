package service

import (
	"context"

	"github.com/Brainsoft-Raxat/respire-api-go/config"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/data"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/models"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/repository"
	"go.uber.org/zap"
)

type UserService interface {
	CreateUser(ctx context.Context, req data.CreateUserRequest) (data.CreateUserResponse, error)
	GetUserByID(ctx context.Context, req data.GetUserByIDRequest) (data.GetUserByIDResponse, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	SearchUsersByUsername(ctx context.Context, req data.SearchUsersByUsernameRequest) (data.SearchUsersByUsernameResponse, error)
	UpdateUser(ctx context.Context, id string, req data.UpdateUserRequest) (data.UpdateUserResponse, error)
	DeleteUser(ctx context.Context, id string) error
}

type FriendshipService interface {
	CreateFriendshipInvitation(ctx context.Context, req data.CreateFriendshipInvitationRequest) (data.CreateFriendshipInvitationResponse, error)
	GetFriendshipInvitations(ctx context.Context, req data.GetFriendshipInvitationsRequest) (data.GetFriendshipInvitationsResponse, error)
	GetFriendsList(ctx context.Context, req data.GetFriendsListRequest) (data.GetFriendsListResponse, error)

	HandleFriendshipInvitation(ctx context.Context, req data.HandleFriendshipInvitationRequest) (data.HandleFriendshipInvitationResponse, error)
	RemoveFriend(ctx context.Context, req data.RemoveFriendRequest) (data.RemoveFriendResponse, error)
	// HandleFriendshipInvitation(ctx context.Context, invitationID string, accept bool) error
	// GetFriendshipInvitations(ctx context.Context) ([]*models.Invitation, error)
	// // GetFriendsList(ctx context.Context) ([]*models.ShortUser, error)
	// DeleteFriendship(ctx context.Context, friendID string) error
}
type SessionService interface {
	CreateSession(ctx context.Context, req data.CreateSessionRequest) (data.CreateSessionResponse, error)
	GetSessionByID(ctx context.Context, req data.GetSessionByIDRequest) (data.GetSessionByIDResponse, error)
	GetSessionByUserID(ctx context.Context, req data.GetSessionByUserIDRequest) (data.GetSessionByUserIDResponse, error)
	UpdateSession(ctx context.Context, id string, req data.UpdateSessionRequest) (data.UpdateSessionResponse, error)
	DeleteSession(ctx context.Context, id string) error
	GetSessionsByUserIDAndDateRange(ctx context.Context, req data.GetSessionByUserIDAndDateRequest) (data.GetSessionByUserIDAndDateResponse, error)
}

type Service struct {
	UserService
	FriendshipService
	SessionService
}

func New(repos *repository.Repository, cfg *config.Configs, logger *zap.SugaredLogger) *Service {
	return &Service{
		UserService:       NewUserService(repos, cfg, logger),
		FriendshipService: NewFriendshipService(repos, cfg, logger),
		SessionService:    NewSessionService(repos, cfg, logger),
	}
}
