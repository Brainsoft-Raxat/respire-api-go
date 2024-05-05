package service

import (
	"context"
	"time"

	"github.com/Brainsoft-Raxat/respire-api-go/config"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/data"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/models"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/repository"
	"github.com/Brainsoft-Raxat/respire-api-go/pkg/ctxconst"
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
	now := time.Now().UTC()
	from := now.AddDate(0, 0, -7)
	if req.ID == "" {
		req.ID = ctxconst.GetUserID(ctx)
	}
	sessions, err := s.SessionRepo.GetSessionsByUserID(ctx, req.ID, from, now, 0)
	if err != nil {
		return data.GetSessionByUserIDResponse{}, err
	}

	return data.GetSessionByUserIDResponse{
		Sum: sessions.Sum(),
	}, nil
}

func (s *sessionService) GetSessionsByUserIDAndDateRange(ctx context.Context, req data.GetSessionByUserIDAndDateRequest) (data.GetSessionByUserIDAndDateResponse, error) {
	if req.ID == "" {
		req.ID = ctxconst.GetUserID(ctx)
	}

	var timeRange [2]time.Time
	now := time.Now()
	switch req.Period {
	case "week":
		timeRange = repository.GetWeek(now)
	case "month":
		timeRange = repository.GetMonth(now)
	default:
		timeRange[0] = req.Start
		timeRange[1] = req.End
	}
	sessions, err := s.SessionRepo.GetSessionsByUserIDAndDateRange(ctx, req.ID, timeRange)
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
	if req.UID == "" {
		req.UID = ctxconst.GetUserID(ctx)
	}
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

func (s *sessionService) GetUserStat(ctx context.Context, req data.GetUserStatRequest) (data.GetUserStatResponse, error) {
	if req.ID == "" {
		req.ID = ctxconst.GetUserID(ctx)
	}
	var err error
	resp := data.GetUserStatResponse{}
	resp.CurrentStreak, resp.BiggestStreak, resp.SavedMoney, err = s.SessionRepo.GetUserStat(ctx, req.ID)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
