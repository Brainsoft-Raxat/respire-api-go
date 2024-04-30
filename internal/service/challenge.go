package service

import (
	"context"

	"github.com/Brainsoft-Raxat/respire-api-go/config"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/data"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/models"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/repository"
	"github.com/Brainsoft-Raxat/respire-api-go/pkg/ctxconst"
	"go.uber.org/zap"
)

type challengeService struct {
	challengeRepo repository.ChallengeRepository
	cfg           *config.Configs
	log           *zap.SugaredLogger
}

func NewChallengeService(repo *repository.Repository, cfg *config.Configs, log *zap.SugaredLogger) ChallengeService {
	return &challengeService{
		challengeRepo: repo.ChallengeRepository,
		cfg:           cfg,
		log:           log,
	}
}

func (s *challengeService) CreateChallenge(ctx context.Context, req data.CreateChallengeRequest) (data.CreateChallengeResponse, error) {
	// check if user is already in a challenge
	userID := ctxconst.GetUserID(ctx)

	challenge := models.Challenge{
		Type:            req.Type,
		Name:            req.Name,
		Description:     req.Description,
		EndDate:         req.EndDate,
		OwnerID:         userID,
		Participants:    []string{userID},
		Invited:         req.Invited,
		Prize:           req.Prize,
		Penalty:         req.Penalty,
		CigarettesLimit: req.CigarettesLimit,
	}

	id, err := s.challengeRepo.CreateChallenge(ctx, &challenge)
	if err != nil {
		return data.CreateChallengeResponse{}, err
	}

	return data.CreateChallengeResponse{
		ID: id,
	}, nil
}

func (s *challengeService) GetChallengeByID(ctx context.Context, req data.GetChallengeByIDRequest) (data.GetChallengeByIDResponse, error) {
	challenge, err := s.challengeRepo.GetChallengeByID(ctx, req.ID)
	if err != nil {
		return data.GetChallengeByIDResponse{}, err
	}

	return data.GetChallengeByIDResponse{
		Challenge: challenge,
	}, nil
}

func (s *challengeService) GetChallengesByUserID(ctx context.Context, req data.GetChallengesByUserIDRequest) (data.GetChallengesByUserIDResponse, error) {
	if req.UserID == "" {
		req.UserID = ctxconst.GetUserID(ctx)
	}

	challenges, err := s.challengeRepo.GetChallengesByUserID(ctx, req.UserID, req.ChallengeType, req.Invite, req.Limit, req.Page, req.From, req.To)
	if err != nil {
		return data.GetChallengesByUserIDResponse{}, err
	}

	return data.GetChallengesByUserIDResponse{
		Challenges: challenges,
	}, nil
}

func (s *challengeService) UpdateChallengeByID(ctx context.Context, req data.UpdateChallengeRequest) (data.UpdateChallengeResponse, error) {
	challenge := models.Challenge{
		Name:            req.Name,
		Description:     req.Description,
		EndDate:         req.EndDate,
		Invited:         req.Invited,
		Prize:           req.Prize,
		Penalty:         req.Penalty,
		CigarettesLimit: req.CigarettesLimit,
	}

	err := s.challengeRepo.UpdateChallenge(ctx, req.ID, &challenge)
	if err != nil {
		return data.UpdateChallengeResponse{}, err
	}

	return data.UpdateChallengeResponse{
		ID: req.ID,
	}, nil
}

func (s *challengeService) DeleteChallengeByID(ctx context.Context, req data.GetChallengeByIDRequest) error {
	err := s.challengeRepo.DeleteChallenge(ctx, req.ID)
	if err != nil {
		return err
	}

	return nil
}