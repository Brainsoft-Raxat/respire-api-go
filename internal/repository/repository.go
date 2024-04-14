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
	FriendsGroupRepository
	ChallengeRepository
	TaskRepository
}

type UserRepository interface {
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) (string, error)
	UpdateUser(ctx context.Context, id string, user *models.User) error
	DeleteUser(ctx context.Context, id string) error
}

type InvitationRepository interface {
	GetInvitationsByUserID(ctx context.Context, userID string) ([]*models.Invitation, error)
	GetInvitationByID(ctx context.Context, id string) (*models.Invitation, error)
	CreateInvitation(ctx context.Context, invitation *models.Invitation) error
	UpdateInvitation(ctx context.Context, id string, invitation *models.Invitation) error
	DeleteInvitation(ctx context.Context, id string) error
}

type FriendshipRepository interface {
	GetFriendshipsByUserID(ctx context.Context, userID string) ([]*models.Friendship, error)
	GetFriendshipByID(ctx context.Context, id string) (*models.Friendship, error)
	CreateFriendship(ctx context.Context, friendship *models.Friendship) error
	UpdateFriendship(ctx context.Context, id string, friendship *models.Friendship) error
	DeleteFriendship(ctx context.Context, id string) error
	AreFriends(ctx context.Context, userID1 string, userID2 string) (bool, error)
}

type FriendsGroupRepository interface {
	GetFriendsGroupByID(ctx context.Context, id string) (*models.FriendsGroup, error)
	GetFriendsGroupsByUserID(ctx context.Context, userID string) ([]*models.FriendsGroup, error)
	CreateFriendsGroup(ctx context.Context, friendsGroup *models.FriendsGroup) error
	UpdateFriendsGroup(ctx context.Context, id string, friendsGroup *models.FriendsGroup) error
	DeleteFriendsGroup(ctx context.Context, id string) error

	CreateFriendsGroupInvitation(ctx context.Context, invitation *models.FriendsGroupInvitation) error
	CreateMultipleFriendsGroupInvitations(ctx context.Context, invitations []*models.FriendsGroupInvitation) error
	DeleteFriendsGroupInvitation(ctx context.Context, id string) error
	GetFriendsGroupInvitationByID(ctx context.Context, id string) (*models.FriendsGroupInvitation, error)
	GetFriendsGroupInvitationsByUserID(ctx context.Context, userID string) ([]*models.FriendsGroupInvitation, error)
	DeleteMultipleFriendsGroupInvitationsByFriendGroupID(ctx context.Context, friendGroupID string) error
}

type ChallengeRepository interface {
	GetChallengeByID(ctx context.Context, friendsGroupID, id string) (*models.Challenge, error)
	GetChallengesByFriendsGroupID(ctx context.Context, friendsGroupID string, filters map[string]interface{}, date *time.Time) ([]*models.Challenge, error)
	CreateChallenge(ctx context.Context, friendsGroupID string, challenge *models.Challenge) error
	UpdateChallenge(ctx context.Context, friendsGroupID, id string, challenge *models.Challenge) error
}

type TaskRepository interface {
	GetTaskByID(ctx context.Context, friendsGroupID, challengeID, id string) (*models.Task, error)
	GetTasksByChallengeID(ctx context.Context, friendsGroupID, challengeID string) ([]*models.Task, error)
	CreateTask(ctx context.Context, friendsGroupID, challengeID string, task *models.Task) error
	UpdateTask(ctx context.Context, friendsGroupID, challengeID, id string, task *models.Task) error
	DeleteTask(ctx context.Context, friendsGroupID, challengeID, id string) error
}

func New(conn *connection.Connection, cfg *config.Configs, logger *zap.SugaredLogger) *Repository {
	return &Repository{
		UserRepository:         NewUserRepository(conn.Firestore, cfg.Firebase, logger),
		InvitationRepository:   NewInvitationRepository(conn.Firestore, cfg.Firebase, logger),
		FriendshipRepository:   NewFriendshipRepository(conn.Firestore, cfg.Firebase, logger),
		FriendsGroupRepository: NewFriendsGroupRepository(conn.Firestore, cfg.Firebase, logger),
		ChallengeRepository:    NewChallengeRepository(conn.Firestore, cfg.Firebase, logger),
		TaskRepository:         NewTaskRepository(conn.Firestore, cfg.Firebase, logger),
	}
}
