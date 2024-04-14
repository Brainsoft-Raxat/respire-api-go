package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"

	"github.com/Brainsoft-Raxat/respire-api-go/config"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/models"
	"go.uber.org/zap"
)

const (
	friendshipCollection = "friendships"
)

type friendshipRepository struct {
	client *firestore.Client
	cfg    *config.Firebase
	logger *zap.SugaredLogger
}

func NewFriendshipRepository(client *firestore.Client, cfg *config.Firebase, logger *zap.SugaredLogger) FriendshipRepository {
	return &friendshipRepository{
		client: client,
		cfg:    cfg,
		logger: logger,
	}
}

func (r *friendshipRepository) GetFriendshipsByUserID(ctx context.Context, userID string) ([]*models.Friendship, error) {
	query := r.client.Collection(friendshipCollection).Where("users", "array-contains", userID)
	iter := query.Documents(ctx)
	var friendships []*models.Friendship
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			r.logger.Errorf("Failed to iterate friendship documents: %v", err)
			return nil, err
		}

		var friendship models.Friendship
		if err := doc.DataTo(&friendship); err != nil {
			r.logger.Errorf("Failed to convert friendship data to struct: %v", err)
			return nil, err
		}

		friendship.ID = doc.Ref.ID
		friendships = append(friendships, &friendship)
	}

	return friendships, nil
}

func (r *friendshipRepository) CreateFriendship(ctx context.Context, friendship *models.Friendship) error {
	_, err := r.client.Collection(friendshipCollection).Doc(friendship.ID).Set(ctx, friendship)
	if err != nil {
		r.logger.Errorf("Failed to create friendship: %v", err)
		return err
	}
	return nil
}

func (r *friendshipRepository) DeleteFriendship(ctx context.Context, id string) error {
	_, err := r.client.Collection(friendshipCollection).Doc(id).Delete(ctx)
	if err != nil {
		r.logger.Errorf("Failed to delete friendship: %v", err)
		return err
	}
	return nil
}

func (r *friendshipRepository) GetFriendshipByID(ctx context.Context, id string) (*models.Friendship, error) {
	doc, err := r.client.Collection(friendshipCollection).Doc(id).Get(ctx)
	if err != nil {
		r.logger.Errorf("Failed to get friendship by ID: %v", err)
		return nil, err
	}

	var friendship models.Friendship
	if err := doc.DataTo(&friendship); err != nil {
		r.logger.Errorf("Failed to convert friendship data to struct: %v", err)
		return nil, err
	}

	friendship.ID = doc.Ref.ID

	return &friendship, nil
}

func (r *friendshipRepository) UpdateFriendship(ctx context.Context, id string, friendship *models.Friendship) error {
	_, err := r.client.Collection(friendshipCollection).Doc(id).Set(ctx, friendship, firestore.MergeAll)
	if err != nil {
		r.logger.Errorf("Failed to update friendship: %v", err)
		return err
	}
	return nil
}

func (r *friendshipRepository) AreFriends(ctx context.Context, userID1 string, userID2 string) (bool, error) {
	query := r.client.Collection(friendshipCollection).Where("users", "array-contains", userID1).Where("users", "array-contains", userID2).Limit(1)
	_, err := query.Documents(ctx).Next()
	if err == iterator.Done {
		return false, nil
	}
	if err != nil {
		r.logger.Errorf("Failed to get friendship document: %v", err)
		return false, err
	}

	return true, nil
}
