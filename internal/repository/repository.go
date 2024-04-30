package repository

import (
	"context"
	"time"

	"github.com/Brainsoft-Raxat/respire-api-go/config"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/app/connection"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/models"
	"go.uber.org/zap"
)

type Repository struct {
	UserRepository
	InvitationRepository
	FriendshipRepository
	SessionRepository
	ChallengeRepository
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

type SessionRepository interface {
	CreateSession(ctx context.Context, user *models.SmokeSession) (string, error)
	GetSessionByID(ctx context.Context, id string) (*models.SmokeSession, error)
	GetSessionsByUserID(ctx context.Context, id string) (SessionsInfo, error)
	GetSessionsByUserIDAndDateRange(ctx context.Context, username string, ti [2]time.Time) (SessionsInfo, error)
	UpdateSession(ctx context.Context, excludeID string, ses *models.SmokeSession) error
	DeleteSession(ctx context.Context, id string) error
}

type ChallengeRepository interface {
	CreateChallenge(ctx context.Context, challenge *models.Challenge) (string, error)
	GetChallengeByID(ctx context.Context, challengeID string) (*models.Challenge, error)
	GetChallengesByUserID(ctx context.Context, userID, challengeType string, invite bool, limit, page int, from, to time.Time) ([]*models.Challenge, error)
	UpdateChallenge(ctx context.Context, id string, challenge *models.Challenge) error
	DeleteChallenge(ctx context.Context, id string) error
}

type TaskRepository interface{}

func New(conn *connection.Connection, cfg *config.Configs, logger *zap.SugaredLogger) *Repository {
	return &Repository{
		UserRepository:       NewUserRepository(conn.Firestore, cfg.Firebase, logger),
		InvitationRepository: NewInvitationRepository(conn.Firestore, cfg.Firebase, logger),
		FriendshipRepository: NewFriendshipRepository(conn.Firestore, cfg.Firebase, logger),
		SessionRepository:    NewSessionRepository(conn.Firestore, cfg.Firebase, logger),
		ChallengeRepository:  NewChallengeRepository(conn.Firestore, cfg.Firebase, logger),
	}
}
