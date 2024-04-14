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

type friendshipService struct {
	repo *repository.Repository
	cfg  *config.Configs
	log  *zap.SugaredLogger
}

func NewFriendshipService(repo *repository.Repository, cfg *config.Configs, log *zap.SugaredLogger) FriendshipService {
	return &friendshipService{
		repo: repo,
		cfg:  cfg,
		log:  log,
	}
}

func (s *friendshipService) CreateFriendshipInvitation(ctx context.Context, friendID string) error {
	userID := ctx.Value("uid").(string)

	areFriends, err := s.repo.FriendshipRepository.AreFriends(ctx, userID, friendID)
	if err != nil {
		return err
	}

	if areFriends {
		return fmt.Errorf("users are already friends")
	}

	err = s.repo.InvitationRepository.CreateInvitation(ctx, &models.Invitation{
		FromUserID: userID,
		ToUserID:   friendID,
		Status:     models.InvitationPending,
		SentDate:   time.Now(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *friendshipService) HandleFriendshipInvitation(ctx context.Context, invitationID string, accept bool) error {
	userID := ctx.Value("uid").(string)

	invitation, err := s.repo.InvitationRepository.GetInvitationByID(ctx, invitationID)
	if err != nil {
		return err
	}

	if invitation.ToUserID != userID {
		return fmt.Errorf("user is not the recipient of the invitation")
	}

	if invitation.Status != models.InvitationPending {
		return fmt.Errorf("invitation is not pending")
	}

	if accept {
		err = s.repo.FriendshipRepository.CreateFriendship(ctx, &models.Friendship{
			Users: []string{invitation.FromUserID, invitation.ToUserID},
			Since: time.Now(),
		})
		if err != nil {
			return err
		}
	}

	err = s.repo.InvitationRepository.DeleteInvitation(ctx, invitationID)
	if err != nil {
		return err
	}

	return nil
}

func (s *friendshipService) GetFriendshipInvitations(ctx context.Context) ([]*models.Invitation, error) {
	userID := ctx.Value("uid").(string)

	invitations, err := s.repo.InvitationRepository.GetInvitationsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	for i, invitation := range invitations {
		sender, err := s.repo.UserRepository.GetUserByID(ctx, invitation.FromUserID)
		if err != nil {
			return nil, err
		}
		invitations[i].FromUserID = sender.ID
	}

	return invitations, nil
}

func (s *friendshipService) GetFriendsList(ctx context.Context) ([]*models.Friend, error) {
	userID := ctx.Value("uid").(string)

	friendships, err := s.repo.FriendshipRepository.GetFriendshipsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var friends []*models.Friend
	for _, friendship := range friendships {
		friendID := friendship.Users[0]
		if friendID == userID {
			friendID = friendship.Users[1]
		}

		friend, err := s.repo.UserRepository.GetUserByID(ctx, friendID)
		if err != nil {
			return nil, err
		}

		friends = append(friends, &models.Friend{
			ID:    friend.ID,
			Name:  friend.Name,
			Since: friendship.Since,
		})
	}

	return friends, nil
}

func (s *friendshipService) DeleteFriendship(ctx context.Context, friendID string) error {
	userID := ctx.Value("uid").(string)

	areFriends, err := s.repo.FriendshipRepository.AreFriends(ctx, userID, friendID)
	if err != nil {
		return err
	}

	if !areFriends {
		return fmt.Errorf("users are not friends")
	}

	friendships, err := s.repo.FriendshipRepository.GetFriendshipsByUserID(ctx, userID)
	if err != nil {
		return err
	}

	var friendshipID string
	for _, friendship := range friendships {
		if (friendship.Users[0] == userID && friendship.Users[1] == friendID) ||
			(friendship.Users[0] == friendID && friendship.Users[1] == userID) {
			friendshipID = friendship.ID
			break
		}
	}

	err = s.repo.FriendshipRepository.DeleteFriendship(ctx, friendshipID)
	if err != nil {
		return err
	}

	return nil
}