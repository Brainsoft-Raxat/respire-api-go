package repository

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/Brainsoft-Raxat/respire-api-go/config"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/models"
	"go.uber.org/zap"
)

const (
	FriendsGroupCollection            = "friends_groups"
	FriendsGroupInvitationsCollection = "friends_group_invitations"
)

type friendsGroupRepository struct {
	client *firestore.Client
	cfg    *config.Firebase
	logger *zap.SugaredLogger
}

func NewFriendsGroupRepository(client *firestore.Client, cfg *config.Firebase, logger *zap.SugaredLogger) FriendsGroupRepository {
	return &friendsGroupRepository{
		client: client,
		cfg:    cfg,
		logger: logger,
	}
}

func (r *friendsGroupRepository) GetFriendsGroupByID(ctx context.Context, id string) (*models.FriendsGroup, error) {
	doc, err := r.client.Collection(FriendsGroupCollection).Doc(id).Get(ctx)
	if err != nil {
		r.logger.Errorf("Failed to get friends group by ID: %v", err)
		return nil, err
	}

	var friendsGroup models.FriendsGroup
	if err := doc.DataTo(&friendsGroup); err != nil {
		r.logger.Errorf("Failed to convert friends group data to struct: %v", err)
		return nil, err
	}

	friendsGroup.ID = doc.Ref.ID
	return &friendsGroup, nil
}

func (r *friendsGroupRepository) GetFriendsGroupsByUserID(ctx context.Context, userID string) ([]*models.FriendsGroup, error) {
	query := r.client.Collection(FriendsGroupCollection).Where("friends", "array-contains", userID)
	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		r.logger.Errorf("Failed to get friends groups by user ID: %v", err)
		return nil, err
	}

	var friendsGroups []*models.FriendsGroup
	for _, doc := range docs {
		var friendsGroup models.FriendsGroup
		if err := doc.DataTo(&friendsGroup); err != nil {
			r.logger.Errorf("Failed to convert friends group data to struct: %v", err)
			return nil, err
		}
		friendsGroup.ID = doc.Ref.ID
		friendsGroups = append(friendsGroups, &friendsGroup)
	}

	return friendsGroups, nil
}

func (r *friendsGroupRepository) CreateFriendsGroup(ctx context.Context, friendsGroup *models.FriendsGroup) error {
	_, _, err := r.client.Collection(FriendsGroupCollection).Add(ctx, friendsGroup)
	if err != nil {
		r.logger.Errorf("Failed to create friends group: %v", err)
		return err
	}
	return nil
}

func (r *friendsGroupRepository) UpdateFriendsGroup(ctx context.Context, id string, friendsGroup *models.FriendsGroup) error {
	_, err := r.client.Collection(FriendsGroupCollection).Doc(id).Set(ctx, friendsGroup)
	if err != nil {
		r.logger.Errorf("Failed to update friends group: %v", err)
		return err
	}
	return nil
}

func (r *friendsGroupRepository) DeleteFriendsGroup(ctx context.Context, id string) error {
	_, err := r.client.Collection(FriendsGroupCollection).Doc(id).Delete(ctx)
	if err != nil {
		r.logger.Errorf("Failed to delete friends group: %v", err)
		return err
	}
	return nil
}

func (r *friendsGroupRepository) CreateFriendsGroupInvitation(ctx context.Context, invitation *models.FriendsGroupInvitation) error {
	_, _, err := r.client.Collection(FriendsGroupInvitationsCollection).Add(ctx, invitation)
	if err != nil {
		r.logger.Errorf("Failed to create friends group invitation: %v", err)
		return err
	}
	return nil
}

func (r *friendsGroupRepository) CreateMultipleFriendsGroupInvitations(ctx context.Context, invitations []*models.FriendsGroupInvitation) error {
	bulkWriter := r.client.BulkWriter(ctx)

	var jobs []*firestore.BulkWriterJob
	var errors []error

	for _, invitation := range invitations {
		docRef := r.client.Collection(FriendsGroupInvitationsCollection).NewDoc()
		job, err := bulkWriter.Create(docRef, invitation)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		jobs = append(jobs, job)
	}

	bulkWriter.End()

	for _, job := range jobs {
		if result, err := job.Results(); err != nil {
			errors = append(errors, err)
		} else {
			fmt.Printf("Write result: %v\n", result)
		}
	}

	if len(errors) > 0 {
		return errors[0]
	}

	r.logger.Info("All invitations have been enqueued and processed successfully.")
	return nil
}

func (r *friendsGroupRepository) DeleteFriendsGroupInvitation(ctx context.Context, id string) error {
	_, err := r.client.Collection(FriendsGroupInvitationsCollection).Doc(id).Delete(ctx)
	if err != nil {
		r.logger.Errorf("Failed to delete friends group invitation: %v", err)
		return err
	}
	return nil
}

func (r *friendsGroupRepository) GetFriendsGroupInvitationByID(ctx context.Context, id string) (*models.FriendsGroupInvitation, error) {
	doc, err := r.client.Collection(FriendsGroupInvitationsCollection).Doc(id).Get(ctx)
	if err != nil {
		r.logger.Errorf("Failed to get friends group invitation by ID: %v", err)
		return nil, err
	}

	var invitation models.FriendsGroupInvitation
	if err := doc.DataTo(&invitation); err != nil {
		r.logger.Errorf("Failed to convert invitation data to struct: %v", err)
		return nil, err
	}

	invitation.ID = doc.Ref.ID
	return &invitation, nil
}

func (r *friendsGroupRepository) GetFriendsGroupInvitationsByUserID(ctx context.Context, userID string) ([]*models.FriendsGroupInvitation, error) {
	query := r.client.Collection(FriendsGroupInvitationsCollection).Where("to_user_id", "==", userID)
	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		r.logger.Errorf("Failed to get friends group invitations by user ID: %v", err)
		return nil, err
	}

	var invitations []*models.FriendsGroupInvitation
	for _, doc := range docs {
		var invitation models.FriendsGroupInvitation
		if err := doc.DataTo(&invitation); err != nil {
			r.logger.Errorf("Failed to convert invitation data to struct: %v", err)
			return nil, err
		}
		invitation.ID = doc.Ref.ID
		invitations = append(invitations, &invitation)
	}

	return invitations, nil
}

func (s *friendsGroupRepository) DeleteMultipleFriendsGroupInvitationsByFriendGroupID(ctx context.Context, friendGroupID string) error {
	query := s.client.Collection(FriendsGroupInvitationsCollection).Where("friends_group_id", "==", friendGroupID)
	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		s.logger.Errorf("Failed to get friends group invitations by friends group ID: %v", err)
		return err
	}

	for _, doc := range docs {
		_, err := doc.Ref.Delete(ctx)
		if err != nil {
			s.logger.Errorf("Failed to delete friends group invitation: %v", err)
			return err
		}
	}

	return nil
}
