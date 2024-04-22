package service

import (
	"context"
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

type userService struct {
	cfg             *config.Configs
	logger          *zap.SugaredLogger
	userRepo        repository.UserRepository
	friendshipsRepo repository.FriendshipRepository
}

func NewUserService(repo *repository.Repository, cfg *config.Configs, logger *zap.SugaredLogger) UserService {
	return &userService{
		cfg:             cfg,
		logger:          logger,
		userRepo:        repo.UserRepository,
		friendshipsRepo: repo.FriendshipRepository,
	}
}

func (s *userService) GetUserByID(ctx context.Context, req data.GetUserByIDRequest) (data.GetUserByIDResponse, error) {
	if req.ID == "" {
		req.ID = ctxconst.GetUserID(ctx)
	}

	user, err := s.userRepo.GetUserByID(ctx, req.ID)
	if err != nil {
		return data.GetUserByIDResponse{}, err
	}

	friendsCount, err := s.friendshipsRepo.GetFriendshipsCountByUserID(ctx, req.ID)
	if err != nil {
		return data.GetUserByIDResponse{}, err
	}

	return data.GetUserByIDResponse{
		User:    user,
		Friends: friendsCount,
	}, nil
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.userRepo.GetUserByEmail(ctx, email)
}

func (s *userService) DeleteUser(ctx context.Context, id string) error {
	return s.userRepo.DeleteUser(ctx, id)
}

func (s *userService) CreateUser(ctx context.Context, req data.CreateUserRequest) (data.CreateUserResponse, error) {
	// TODO: add checks if user already exists by email etc..
	user := &models.User{
		Name:          req.Name,
		Email:         req.Email,
		LongestStreak: 0,
		Status:        models.STATUS_CREATED,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	id, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return data.CreateUserResponse{}, err
	}

	return data.CreateUserResponse{ID: id}, nil
}

func (s *userService) UpdateUser(ctx context.Context, id string, req data.UpdateUserRequest) (data.UpdateUserResponse, error) {
	user, err := s.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return data.UpdateUserResponse{}, err
	}

	if user.Status == models.STATUS_CREATED {
		if req.Username != "" && !req.QuitDate.IsZero() {
			_, err := s.userRepo.GetUserByUsername(ctx, req.Username)
			if err != nil {
				if apperror.EqualWithErrorCode(err, errcodes.NotFoundError) {
					user.Status = models.STATUS_COMPLETED
				} else {
					return data.UpdateUserResponse{}, apperror.NewErrorInfo(ctx, errcodes.InvalidRequest, "Username already exists or other issue")
				}
			}

			user.Username = req.Username
			user.QuitDate = req.QuitDate
		} else {
			return data.UpdateUserResponse{}, apperror.NewErrorInfo(ctx, errcodes.InvalidRequest, "Username and quit date must be provided")
		}
	} else {
		if user.Username != req.Username {
			_, err := s.userRepo.GetUserByUsername(ctx, req.Username)
			if err != nil {
				if apperror.EqualWithErrorCode(err, errcodes.NotFoundError) {
					user.Username = req.Username
				} else {
					return data.UpdateUserResponse{}, apperror.NewErrorInfo(ctx, errcodes.InvalidRequest, "Username already exists or other issue")
				}
			}

			user.Username = req.Username
		}
	}

	user.Name = req.Name
	user.Avatar = req.Avatar
	user.UpdatedAt = time.Now()

	err = s.userRepo.UpdateUser(ctx, id, user)
	if err != nil {
		return data.UpdateUserResponse{}, err
	}

	return data.UpdateUserResponse{ID: id}, nil
}

func (s *userService) SearchUsersByUsername(ctx context.Context, req data.SearchUsersByUsernameRequest) (data.SearchUsersByUsernameResponse, error) {
	userID := ctxconst.GetUserID(ctx)

	users, err := s.userRepo.SearchUsersByUsernamePrefix(ctx, userID, req.Username, req.Limit)
	if err != nil {
		return data.SearchUsersByUsernameResponse{}, err
	}

	return data.SearchUsersByUsernameResponse{Users: users}, nil
}
