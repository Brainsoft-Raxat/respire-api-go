package service

import (
	"context"
	"net/http"
	"time"

	"github.com/Brainsoft-Raxat/respire-api-go/config"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/data"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/models"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/repository"
	"github.com/Brainsoft-Raxat/respire-api-go/pkg/apperror"
	"github.com/Brainsoft-Raxat/respire-api-go/pkg/ctxconst"
	"github.com/Brainsoft-Raxat/respire-api-go/pkg/errcodes"
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

func (s *friendshipService) CreateFriendshipInvitation(ctx context.Context, req data.CreateFriendshipInvitationRequest) (data.CreateFriendshipInvitationResponse, error) {
	userID := ctxconst.GetUserID(ctx)

	areFriends, err := s.repo.FriendshipRepository.AreFriends(ctx, userID, req.FriendID)
	if err != nil {
		return data.CreateFriendshipInvitationResponse{}, err
	}

	if areFriends {
		return data.CreateFriendshipInvitationResponse{}, apperror.NewErrorInfo(ctx, errcodes.UserAreAlreadyFriends, "users are already friends")
	}

	err = s.repo.InvitationRepository.CreateInvitation(ctx, &models.Invitation{
		FromUserID: userID,
		ToUserID:   req.FriendID,
		Status:     models.InvitationPending,
		SentDate:   time.Now(),
	})
	if err != nil {
		return data.CreateFriendshipInvitationResponse{}, err
	}

	return data.CreateFriendshipInvitationResponse{
		Status:  http.StatusOK,
		Message: "Friend invitation sent successfully",
	}, nil
}

func (s *friendshipService) GetFriendshipInvitations(ctx context.Context, req data.GetFriendshipInvitationsRequest) (data.GetFriendshipInvitationsResponse, error) {
	userID := ctxconst.GetUserID(ctx)

	invitations, err := s.repo.InvitationRepository.GetInvitationsByUserID(ctx, userID)
	if err != nil {
		return data.GetFriendshipInvitationsResponse{}, err
	}

	return data.GetFriendshipInvitationsResponse{
		Invitations: invitations,
	}, nil
}

func (s *friendshipService) HandleFriendshipInvitation(ctx context.Context, req data.HandleFriendshipInvitationRequest) (data.HandleFriendshipInvitationResponse, error) {
	userID := ctxconst.GetUserID(ctx)

	invitation, err := s.repo.InvitationRepository.GetInvitationByID(ctx, req.InvitationID)
	if err != nil {
		return data.HandleFriendshipInvitationResponse{}, err
	}

	if invitation.ToUserID != userID {
		return data.HandleFriendshipInvitationResponse{}, apperror.NewErrorInfo(ctx, errcodes.InvalidRequest, "user does not have permission")
	}

	if req.Accept {
		err = s.repo.FriendshipRepository.CreateFriendship(ctx, &models.Friendship{
			UserID:   userID,
			FriendID: invitation.FromUserID,
			Since:    time.Now(),
		})
		if err != nil {
			return data.HandleFriendshipInvitationResponse{}, err
		}

		err = s.repo.FriendshipRepository.CreateFriendship(ctx, &models.Friendship{
			UserID:   invitation.FromUserID,
			FriendID: userID,
			Since:    time.Now(),
		})
		if err != nil {
			return data.HandleFriendshipInvitationResponse{}, err
		}
	}

	err = s.repo.InvitationRepository.DeleteInvitation(ctx, req.InvitationID)
	if err != nil {
		return data.HandleFriendshipInvitationResponse{}, err
	}

	return data.HandleFriendshipInvitationResponse{
		Status:  http.StatusOK,
		Message: "Friend invitation handled successfully",
	}, nil
}

func (s *friendshipService) GetFriendsList(ctx context.Context, req data.GetFriendsListRequest) (data.GetFriendsListResponse, error) {
	userID := ctxconst.GetUserID(ctx)

	if req.UserID != userID && req.UserID != "" {
		userID = req.UserID
	}

	friendships, err := s.repo.FriendshipRepository.GetFriendshipsByUserID(ctx, userID)
	if err != nil {
		return data.GetFriendsListResponse{}, err
	}

	var friends []*models.ShortUser

	ids := make([]string, 0, len(friendships))
	for _, friendship := range friendships {
		ids = append(ids, friendship.FriendID)
	}

	if len(ids) > 0 {
		friends, err = s.repo.UserRepository.GetShortUsersByIDs(ctx, ids)
		if err != nil {
			return data.GetFriendsListResponse{}, err
		}
	}

	return data.GetFriendsListResponse{
		Friends: friends,
	}, nil
}

func (s *friendshipService) RemoveFriend(ctx context.Context, req data.RemoveFriendRequest) (data.RemoveFriendResponse, error) {
	userID := ctxconst.GetUserID(ctx)

	areFriends, err := s.repo.FriendshipRepository.AreFriends(ctx, userID, req.FriendID)
	if err != nil {
		return data.RemoveFriendResponse{}, err
	}

	if !areFriends {
		return data.RemoveFriendResponse{}, apperror.NewErrorInfo(ctx, errcodes.Forbidden, "users are already friends")
	}

	err = s.repo.FriendshipRepository.DeleteFriendshipByUsers(ctx, userID, req.FriendID)
	if err != nil {
		return data.RemoveFriendResponse{}, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to remove friend")
	}

	err = s.repo.FriendshipRepository.DeleteFriendshipByUsers(ctx, req.FriendID, userID)
	if err != nil {
		return data.RemoveFriendResponse{}, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to remove friend")
	}

	return data.RemoveFriendResponse{
		Status:  http.StatusOK,
		Message: "Friend removed successfully",
	}, nil
}
