package repository

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/Brainsoft-Raxat/respire-api-go/config"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/models"
	"github.com/Brainsoft-Raxat/respire-api-go/pkg/apperror"
	"github.com/Brainsoft-Raxat/respire-api-go/pkg/errcodes"
	"go.uber.org/zap"
)

const (
	ChallengeCollection = "challenges"
)

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

func (r *challengeRepository) CreateChallenge(ctx context.Context, challenge *models.Challenge) (string, error) {
	ref, _, err := r.client.Collection(ChallengeCollection).Add(ctx, challenge)
	if err != nil {
		r.logger.Errorf("Failed to create challenge: %v", err)
		return "", apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to create challenge")
	}

	return ref.ID, nil
}

func (r *challengeRepository) GetChallengeByID(ctx context.Context, challengeID string) (*models.Challenge, error) {
	doc, err := r.client.Collection(ChallengeCollection).Doc(challengeID).Get(ctx)
	if err != nil {
		r.logger.Errorf("Failed to get challenge by ID: %v", err)
		return nil, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to get challenge by ID")
	}

	var challenge models.Challenge
	if err := doc.DataTo(&challenge); err != nil {
		r.logger.Errorf("Failed to convert challenge data to struct: %v", err)
		return nil, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to get challenge by ID")
	}

	return &challenge, nil
}

func (r *challengeRepository) GetChallengesByUserID(ctx context.Context, userID, challengeType string, invite bool, limit, page int, from, to time.Time) ([]*models.Challenge, error) {
	query := r.client.Collection(ChallengeCollection).Query

	if invite {
		query = query.Where("invited", "array-contains", userID)
	} else {
		query = query.Where("participants", "array-contains", userID)
	}

	if challengeType != "" {
		query = query.Where("type", "==", challengeType)
	}

	if limit > 0 {
		query = query.Limit(limit)
	}

	if page > 0 {
		query = query.Offset((page - 1) * limit)
	}

	if !from.IsZero() {
		query = query.Where("end_date", ">=", from)
	}

	if !to.IsZero() {
		query = query.Where("end_date", "<=", to)
	}

	snapshots, err := query.Documents(ctx).GetAll()
	if err != nil {
		r.logger.Errorf("Failed to get challenges by user ID: %v", err)
		return nil, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to get challenges by user ID")
	}

	challenges := make([]*models.Challenge, 0, len(snapshots))
	for _, snapshot := range snapshots {
		var challenge models.Challenge
		if err := snapshot.DataTo(&challenge); err != nil {
			r.logger.Errorf("Failed to convert challenge data to struct: %v", err)
			return nil, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to get challenges by user ID")
		}
		challenges = append(challenges, &challenge)
	}

	return challenges, nil
}

func (r *challengeRepository) UpdateChallenge(ctx context.Context, challengeID string, challenge *models.Challenge) error {
	model, err := CreateUpdateMap(challenge)
	if err != nil {
		return apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to create update map for challenge")
	}

	_, err = r.client.Collection(ChallengeCollection).Doc(challengeID).Set(ctx, model, firestore.MergeAll)
	if err != nil {
		r.logger.Errorf("Failed to update challenge: %v", err)
		return apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to update challenge")
	}
	return nil
}

func (r *challengeRepository) DeleteChallenge(ctx context.Context, challengeID string) error {
	_, err := r.client.Collection(ChallengeCollection).Doc(challengeID).Delete(ctx)
	if err != nil {
		r.logger.Errorf("Failed to delete challenge: %v", err)
		return apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to delete challenge")
	}
	return nil
}
