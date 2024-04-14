package repository

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/Brainsoft-Raxat/respire-api-go/config"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/models"
	"go.uber.org/zap"
)

const ChallengesCollection = "friends_groups"

type challengeRepository struct {
	client *firestore.Client
	cfg    *config.Firebase
	logger *zap.SugaredLogger
}

func NewChallengeRepository(client *firestore.Client, cfg *config.Firebase, logger *zap.SugaredLogger) ChallengeRepository {
	return &challengeRepository{
		client: client,
		cfg:    cfg,
		logger: logger,
	}
}

func (r *challengeRepository) GetChallengeByID(ctx context.Context, friendsGroupID, id string) (*models.Challenge, error) {
	doc, err := r.client.Collection(FriendsGroupCollection).Doc(friendsGroupID).Collection(ChallengesCollection).Doc(id).Get(ctx)
	if err != nil {
		r.logger.Errorf("Failed to get challenge by ID: %v", err)
		return nil, err
	}

	var challenge models.Challenge
	if err := doc.DataTo(&challenge); err != nil {
		r.logger.Errorf("Failed to convert challenge data to struct: %v", err)
		return nil, err
	}

	challenge.ID = doc.Ref.ID
	return &challenge, nil
}

func (r *challengeRepository) GetChallengesByFriendsGroupID(ctx context.Context, friendsGroupID string, filters map[string]interface{}, date *time.Time) ([]*models.Challenge, error) {
	challenges := make([]*models.Challenge, 0)

	query := r.client.Collection(FriendsGroupCollection).Doc(friendsGroupID).Collection(ChallengesCollection).Query
	for field, value := range filters {
		query = query.Where(field, "==", value)
	}

	if date != nil {
		query = query.Where("start_date", "<=", date).Where("end_date", ">=", date)
	}

	snapshot, err := query.Documents(ctx).GetAll()
	if err != nil {
		r.logger.Errorf("Failed to get challenges by friends group ID: %v", err)
		return nil, err
	}

	for _, doc := range snapshot {
		var challenge models.Challenge
		if err := doc.DataTo(&challenge); err != nil {
			r.logger.Errorf("Failed to convert challenge data to struct: %v", err)
			return nil, err
		}
		challenge.ID = doc.Ref.ID
		challenges = append(challenges, &challenge)
	}

	return challenges, nil
}

func (r *challengeRepository) CreateChallenge(ctx context.Context, friendsGroupID string, challenge *models.Challenge) error {
	_, _, err := r.client.Collection(FriendsGroupCollection).Doc(friendsGroupID).Collection(ChallengesCollection).Add(ctx, challenge)
	if err != nil {
		r.logger.Errorf("Failed to create challenge: %v", err)
		return err
	}

	return nil
}

func (r *challengeRepository) UpdateChallenge(ctx context.Context, friendsGroupID, id string, challenge *models.Challenge) error {
	_, err := r.client.Collection(FriendsGroupCollection).Doc(friendsGroupID).Collection(ChallengesCollection).Doc(id).Set(ctx, challenge)
	if err != nil {
		r.logger.Errorf("Failed to update challenge: %v", err)
		return err
	}

	return nil
}

func (r *challengeRepository) DeleteChallenge(ctx context.Context, friendsGroupID, id string) error {
	_, err := r.client.Collection(FriendsGroupCollection).Doc(friendsGroupID).Collection(ChallengesCollection).Doc(id).Delete(ctx)
	if err != nil {
		r.logger.Errorf("Failed to delete challenge: %v", err)
		return err
	}

	return nil
}
