package service

import (
	"context"
	"time"

	"github.com/Brainsoft-Raxat/respire-api-go/config"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/data"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/models"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/repository"
	"go.uber.org/zap"
)

type userService struct {
	cfg      *config.Configs
	logger   *zap.SugaredLogger
	userRepo repository.UserRepository
}

func NewUserService(repo *repository.Repository, cfg *config.Configs, logger *zap.SugaredLogger) UserService {
	return &userService{
		cfg:      cfg,
		logger:   logger,
		userRepo: repo,
	}
}

func (s *userService) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	return s.userRepo.GetUserByID(ctx, id)
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
		CurrentStreak: 0,
		LongestStreak: 0,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	id, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return data.CreateUserResponse{}, err
	}

	return data.CreateUserResponse{ID: id}, nil
}

func (s *userService) UpdateUser(ctx context.Context, id string, user *models.User) error {
	return s.userRepo.UpdateUser(ctx, id, user)
}
