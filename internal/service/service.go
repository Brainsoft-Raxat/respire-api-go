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
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UpdateUser(ctx context.Context, id string, user *models.User) error
	DeleteUser(ctx context.Context, id string) error
}

type FriendshipService interface {
	CreateFriendshipInvitation(ctx context.Context, friendID string) error
	HandleFriendshipInvitation(ctx context.Context, invitationID string, accept bool) error
	GetFriendshipInvitations(ctx context.Context) ([]*models.Invitation, error)
	GetFriendsList(ctx context.Context) ([]*models.Friend, error)
	DeleteFriendship(ctx context.Context, friendID string) error
}

type FriendsGroupService interface {
	CreateFriendsGroup(ctx context.Context, req models.CreateFriendsGroupRequest) error
	GetFriendsGroups(ctx context.Context) ([]*models.FriendsGroup, error)
	GetFriendsGroupByID(ctx context.Context, id string) (*models.FriendsGroup, error)
	UpdateFriendsGroup(ctx context.Context, id string, req models.UpdateFriendsGroupRequest) error
	DeleteFriendsGroup(ctx context.Context, id string) error

	GetFriendsGroupInvitations(ctx context.Context) ([]*models.FriendsGroupInvitation, error)
	HandleFriendsGroupInvitation(ctx context.Context, id string, accept bool) error
}

type Service struct {
	UserService
	FriendshipService
	FriendsGroupService
}

func New(repos *repository.Repository, cfg *config.Configs, logger *zap.SugaredLogger) *Service {
	return &Service{
		UserService:         NewUserService(repos, cfg, logger),
		FriendshipService:   NewFriendshipService(repos, cfg, logger),
		FriendsGroupService: NewFriendsGroupService(repos, cfg, logger),
	}
}
