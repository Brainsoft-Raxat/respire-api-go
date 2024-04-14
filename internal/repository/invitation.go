package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/Brainsoft-Raxat/respire-api-go/config"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/models"
	"go.uber.org/zap"
	"google.golang.org/api/iterator"
)

const (
	invitationCollection = "invitations"
)

type invitationRepository struct {
	client *firestore.Client
	cfg    *config.Firebase
	logger *zap.SugaredLogger
}

func NewInvitationRepository(client *firestore.Client, cfg *config.Firebase, logger *zap.SugaredLogger) InvitationRepository {
	return &invitationRepository{
		client: client,
		cfg:    cfg,
		logger: logger,
	}
}

func (r *invitationRepository) GetInvitationsByUserID(ctx context.Context, userID string) ([]*models.Invitation, error) {
	iter := r.client.Collection(invitationCollection).Where("to_user_id", "==", userID).Documents(ctx)
	var invitations []*models.Invitation
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			r.logger.Errorf("Failed to iterate invitation documents: %v", err)
			return nil, err
		}

		var invitation models.Invitation
		if err := doc.DataTo(&invitation); err != nil {
			r.logger.Errorf("Failed to convert invitation data to struct: %v", err)
			return nil, err
		}

		invitation.ID = doc.Ref.ID
		invitations = append(invitations, &invitation)
	}

	return invitations, nil
}

func (r *invitationRepository) GetInvitationByID(ctx context.Context, id string) (*models.Invitation, error) {
	doc, err := r.client.Collection(invitationCollection).Doc(id).Get(ctx)
	if err != nil {
		r.logger.Errorf("Failed to get invitation document: %v", err)
		return nil, err
	}

	var invitation models.Invitation
	if err := doc.DataTo(&invitation); err != nil {
		r.logger.Errorf("Failed to convert invitation data to struct: %v", err)
		return nil, err
	}

	invitation.ID = doc.Ref.ID

	return &invitation, nil
}

func (r *invitationRepository) CreateInvitation(ctx context.Context, invitation *models.Invitation) error {
	_, _, err := r.client.Collection(invitationCollection).Add(ctx, invitation)
	if err != nil {
		r.logger.Errorf("Failed to create invitation: %v", err)
		return err
	}

	return nil
}

func (r *invitationRepository) UpdateInvitation(ctx context.Context, id string, invitation *models.Invitation) error {
	_, err := r.client.Collection(invitationCollection).Doc(id).Set(ctx, invitation)
	if err != nil {
		r.logger.Errorf("Failed to update invitation: %v", err)
		return err
	}

	return nil
}

func (r *invitationRepository) DeleteInvitation(ctx context.Context, id string) error {
	_, err := r.client.Collection(invitationCollection).Doc(id).Delete(ctx)
	if err != nil {
		r.logger.Errorf("Failed to delete invitation: %v", err)
		return err
	}

	return nil
}
