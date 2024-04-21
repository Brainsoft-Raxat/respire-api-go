package repository

import (
	"context"

	"github.com/Brainsoft-Raxat/respire-api-go/config"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/app/connection"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/models"
	"go.uber.org/zap"
)

type Repository struct {
	UserRepository
	InvitationRepository
	FriendshipRepository
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (string, error)
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	SearchUsersByUsernamePrefix(ctx context.Context, excludeID string, usernamePrefix string, limit int) ([]*models.ShortUser, error)
	GetShortUsersByIDs(ctx context.Context, ids []string) ([]*models.ShortUser, error)
	UpdateUser(ctx context.Context, id string, user *models.User) error
	DeleteUser(ctx context.Context, id string) error
}

type InvitationRepository interface {
	CreateInvitation(ctx context.Context, invitation *models.Invitation) error
	GetInvitationsByUserID(ctx context.Context, userID string) ([]*models.Invitation, error)
	GetInvitationByID(ctx context.Context, id string) (*models.Invitation, error)
	UpdateInvitation(ctx context.Context, id string, invitation *models.Invitation) error
	DeleteInvitation(ctx context.Context, id string) error
}

type FriendshipRepository interface {
	CreateFriendship(ctx context.Context, friendship *models.Friendship) error
	GetFriendshipsCountByUserID(ctx context.Context, userID string) (int, error)
	GetFriendshipsByUserID(ctx context.Context, userID string) ([]*models.Friendship, error)
	GetFriendshipByID(ctx context.Context, id string) (*models.Friendship, error)
	UpdateFriendship(ctx context.Context, id string, friendship *models.Friendship) error
	DeleteFriendship(ctx context.Context, id string) error
	DeleteFriendshipByUsers(ctx context.Context, userID string, friendID string) error

	AreFriends(ctx context.Context, userID1 string, userID2 string) (bool, error)
}

type ChallengeRepository interface{}

type TaskRepository interface{}

func New(conn *connection.Connection, cfg *config.Configs, logger *zap.SugaredLogger) *Repository {
	return &Repository{
		UserRepository:       NewUserRepository(conn.Firestore, cfg.Firebase, logger),
		InvitationRepository: NewInvitationRepository(conn.Firestore, cfg.Firebase, logger),
		FriendshipRepository: NewFriendshipRepository(conn.Firestore, cfg.Firebase, logger),
	}
}
