package service

import (
	"context"

	"github.com/Brainsoft-Raxat/respire-api-go/config"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/data"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/models"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/repository"
	"go.uber.org/zap"
)

type sessionService struct {
	cfg             *config.Configs
	logger          *zap.SugaredLogger
	SessionRepo     repository.SessionRepository
	friendshipsRepo repository.FriendshipRepository
}

func NewSessionService(repo *repository.Repository, cfg *config.Configs, logger *zap.SugaredLogger) *sessionService {
	return &sessionService{
		cfg:             cfg,
		logger:          logger,
		SessionRepo:     repo.SessionRepository,
		friendshipsRepo: repo.FriendshipRepository,
	}
}

func (s *sessionService) GetSessionByID(ctx context.Context, req data.GetSessionByIDRequest) (data.GetSessionByIDResponse, error) {
	Session, err := s.SessionRepo.GetSessionByID(ctx, req.ID)
	if err != nil {
		return data.GetSessionByIDResponse{}, err
	}

	return data.GetSessionByIDResponse{
		Session: Session,
	}, nil
}

func (s *sessionService) GetSessionByUserID(ctx context.Context, req data.GetSessionByUserIDRequest) (data.GetSessionByUserIDResponse, error) {
	sessions, err := s.SessionRepo.GetSessionsByUserID(ctx, req.ID)
	if err != nil {
		return data.GetSessionByUserIDResponse{}, err
	}

	return data.GetSessionByUserIDResponse{
		Sum: sessions.Sum(),
	}, nil
}

func (s *sessionService) GetSessionsByUserIDAndDateRange(ctx context.Context, req data.GetSessionByUserIDAndDateRequest) (data.GetSessionByUserIDAndDateResponse, error) {
	sessions, err := s.SessionRepo.GetSessionsByUserIDAndDateRange(ctx, req.ID, req.DR)
	if err != nil {
		return data.GetSessionByUserIDAndDateResponse{}, err
	}

	return data.GetSessionByUserIDAndDateResponse{
		Sum: sessions.Sum(),
	}, nil
}

func (s *sessionService) DeleteSession(ctx context.Context, id string) error {
	return s.SessionRepo.DeleteSession(ctx, id)
}

func (s *sessionService) CreateSession(ctx context.Context, req data.CreateSessionRequest) (data.CreateSessionResponse, error) {
	// TODO: add checks if Session already exists by email etc..
	Session := &models.SmokeSession{
		UID:       req.UID,
		Count:     req.Count,
		Timestamp: req.Timestamp,
	}

	id, err := s.SessionRepo.CreateSession(ctx, Session)
	if err != nil {
		return data.CreateSessionResponse{}, err
	}

	return data.CreateSessionResponse{ID: id}, nil
}

func (s *sessionService) UpdateSession(ctx context.Context, id string, req data.UpdateSessionRequest) (data.UpdateSessionResponse, error) {
	Session, err := s.SessionRepo.GetSessionByID(ctx, id)
	if err != nil {
		return data.UpdateSessionResponse{}, err
	}

	return data.UpdateSessionResponse{ID: Session.ID}, nil
}