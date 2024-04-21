package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/firestore/apiv1/firestorepb"
	"google.golang.org/api/iterator"

	"github.com/Brainsoft-Raxat/respire-api-go/config"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/models"
	"github.com/Brainsoft-Raxat/respire-api-go/pkg/apperror"
	"github.com/Brainsoft-Raxat/respire-api-go/pkg/errcodes"
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

func (r *friendshipRepository) GetFriendshipsCountByUserID(ctx context.Context, userID string) (int, error) {
	query := r.client.Collection(friendshipCollection).Where("user_id", "==", userID)
	aggregationQuery := query.NewAggregationQuery().WithCount("all")
	results, err := aggregationQuery.Get(ctx)
	if err != nil {
		r.logger.Errorf("Failed to get friendship count: %v", err)
		return 0, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to get friendship count")
	}

	count, ok := results["all"]
	if !ok {
		return 0, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to get friendship count")
	}

	countValue := count.(*firestorepb.Value)

	return int(countValue.GetIntegerValue()), nil
}

func (r *friendshipRepository) GetFriendshipsByUserID(ctx context.Context, userID string) ([]*models.Friendship, error) {
	query := r.client.Collection(friendshipCollection).Where("user_id", "==", userID)
	iter := query.Documents(ctx)
	var friendships []*models.Friendship
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			r.logger.Errorf("Failed to iterate friendship documents: %v", err)
			return nil, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to iterate friendship documents")
		}

		var friendship models.Friendship
		if err := doc.DataTo(&friendship); err != nil {
			r.logger.Errorf("Failed to convert friendship data to struct: %v", err)
			return nil, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to convert friendship data to struct")
		}

		friendship.ID = doc.Ref.ID
		friendships = append(friendships, &friendship)
	}

	return friendships, nil
}

func (r *friendshipRepository) CreateFriendship(ctx context.Context, friendship *models.Friendship) error {
	_, _, err := r.client.Collection(friendshipCollection).Add(ctx, friendship)
	if err != nil {
		r.logger.Errorf("Failed to create friendship: %v", err)
		return apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to create friendship")
	}
	return nil
}

func (r *friendshipRepository) DeleteFriendship(ctx context.Context, id string) error {
	_, err := r.client.Collection(friendshipCollection).Doc(id).Delete(ctx)
	if err != nil {
		r.logger.Errorf("Failed to delete friendship: %v", err)
		return apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to delete friendship")
	}
	return nil
}

func (r *friendshipRepository) GetFriendshipByID(ctx context.Context, id string) (*models.Friendship, error) {
	doc, err := r.client.Collection(friendshipCollection).Doc(id).Get(ctx)
	if err != nil {
		r.logger.Errorf("Failed to get friendship by ID: %v", err)
		return nil, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to get friendship by ID")
	}

	var friendship models.Friendship
	if err := doc.DataTo(&friendship); err != nil {
		r.logger.Errorf("Failed to convert friendship data to struct: %v", err)
		return nil, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to convert friendship data to struct")
	}

	friendship.ID = doc.Ref.ID

	return &friendship, nil
}

func (r *friendshipRepository) UpdateFriendship(ctx context.Context, id string, friendship *models.Friendship) error {
	_, err := r.client.Collection(friendshipCollection).Doc(id).Set(ctx, friendship, firestore.MergeAll)
	if err != nil {
		r.logger.Errorf("Failed to update friendship: %v", err)
		return apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to update friendship")
	}
	return nil
}

func (r *friendshipRepository) AreFriends(ctx context.Context, userID string, friendID string) (bool, error) {
	query := r.client.Collection(friendshipCollection).Where("user_id", "==", userID).Where("friend_id", "==", friendID).Limit(1)
	_, err := query.Documents(ctx).Next()
	if err == iterator.Done {
		return false, nil
	}
	if err != nil {
		r.logger.Errorf("Failed to get friendship document: %v", err)
		return false, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to get friendship document")
	}

	return true, nil
}

func (r *friendshipRepository) DeleteFriendshipByUsers(ctx context.Context, userID string, friendID string) error {
	query := r.client.Collection(friendshipCollection).Where("user_id", "==", userID).Where("friend_id", "==", friendID).Limit(1)
	iter := query.Documents(ctx)
	doc, err := iter.Next()
	if err == iterator.Done {
		return nil
	}
	if err != nil {
		r.logger.Errorf("Failed to get friendship document: %v", err)
		return apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to get friendship document")
	}

	_, err = doc.Ref.Delete(ctx)
	if err != nil {
		r.logger.Errorf("Failed to delete friendship: %v", err)
		return apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to delete friendship")
	}

	return nil
}