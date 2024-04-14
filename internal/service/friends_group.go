package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Brainsoft-Raxat/respire-api-go/config"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/models"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/repository"
	"go.uber.org/zap"
)

type friendsGroupService struct {
	repo *repository.Repository
	cfg  *config.Configs
	log  *zap.SugaredLogger
}

func NewFriendsGroupService(repo *repository.Repository, cfg *config.Configs, log *zap.SugaredLogger) FriendsGroupService {
	return &friendsGroupService{
		repo: repo,
		cfg:  cfg,
		log:  log,
	}
}

func (s *friendsGroupService) CreateFriendsGroup(ctx context.Context, req models.CreateFriendsGroupRequest) error {
	ownerID := ctx.Value("uid").(string)

	friendsGroup := &models.FriendsGroup{
		Name:       req.Name,
		OwnerID:    ownerID,
		Friends:    []string{"ownerID"},
	}

	err := s.repo.FriendsGroupRepository.CreateFriendsGroup(ctx, friendsGroup)
	if err != nil {
		return err
	}

	var invitations []*models.FriendsGroupInvitation
	for _, friendID := range req.Invites {
		invitation := &models.FriendsGroupInvitation{
			FriendsGroupID: friendID,
			FromUserID:     ownerID,
			ToUserID:       friendID,
			Status:         models.InvitationPending,
			SentDate:       time.Now(),
		}
		invitations = append(invitations, invitation)
	}

	err = s.repo.FriendsGroupRepository.CreateMultipleFriendsGroupInvitations(ctx, invitations)
	if err != nil {
		return err
	}

	return nil
}

func (s *friendsGroupService) GetFriendsGroups(ctx context.Context) ([]*models.FriendsGroup, error) {
	userID := ctx.Value("uid").(string)

	friendsGroups, err := s.repo.FriendsGroupRepository.GetFriendsGroupsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return friendsGroups, nil
}

func (s *friendsGroupService) GetFriendsGroupByID(ctx context.Context, id string) (*models.FriendsGroup, error) {
	friendsGroup, err := s.repo.FriendsGroupRepository.GetFriendsGroupByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return friendsGroup, nil
}

func (s *friendsGroupService) DeleteFriendsGroup(ctx context.Context, id string) error {
	userID := ctx.Value("uid").(string)

	friendsGroup, err := s.repo.FriendsGroupRepository.GetFriendsGroupByID(ctx, id)
	if err != nil {
		return err
	}

	if friendsGroup.OwnerID != userID {
		return fmt.Errorf("user is not the owner of the friends group")
	}

	err = s.repo.FriendsGroupRepository.DeleteMultipleFriendsGroupInvitationsByFriendGroupID(ctx, friendsGroup.ID)
	if err != nil {
		return err
	}

	err = s.repo.FriendsGroupRepository.DeleteFriendsGroup(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// TODO: Implement the UpdateFriendsGroup method
func (s *friendsGroupService) UpdateFriendsGroup(ctx context.Context, id string, req models.UpdateFriendsGroupRequest) error {
	userID := ctx.Value("uid").(string)

	friendsGroup, err := s.repo.FriendsGroupRepository.GetFriendsGroupByID(ctx, id)
	if err != nil {
		return err
	}

	if friendsGroup.OwnerID != userID {
		return fmt.Errorf("user is not the owner of the friends group")
	}

	friendsGroup.Name = req.Name

	err = s.repo.FriendsGroupRepository.UpdateFriendsGroup(ctx, id, friendsGroup)
	if err != nil {
		return err
	}

	return nil
}

func (s *friendsGroupService) GetFriendsGroupInvitations(ctx context.Context) ([]*models.FriendsGroupInvitation, error) {
	userID := ctx.Value("uid").(string)

	invitations, err := s.repo.FriendsGroupRepository.GetFriendsGroupInvitationsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return invitations, nil
}

func (s *friendsGroupService) HandleFriendsGroupInvitation(ctx context.Context, id string, accept bool) error {
	userID := ctx.Value("uid").(string)

	invitation, err := s.repo.FriendsGroupRepository.GetFriendsGroupInvitationByID(ctx, id)
	if err != nil {
		return err
	}

	if invitation.ToUserID != userID {
		return fmt.Errorf("user is not the recipient of the invitation")
	}

	if accept {
		friendsGroup, err := s.repo.FriendsGroupRepository.GetFriendsGroupByID(ctx, invitation.FriendsGroupID)
		if err != nil {
			return err
		}

		friendsGroup.Friends = append(friendsGroup.Friends, userID)

		err = s.repo.FriendsGroupRepository.UpdateFriendsGroup(ctx, friendsGroup.ID, friendsGroup)
		if err != nil {
			return err
		}
	}

	err = s.repo.FriendsGroupRepository.DeleteFriendsGroupInvitation(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
