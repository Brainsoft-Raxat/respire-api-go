package service

import (
	"context"
	"slices"
	"strings"
	"time"

	"github.com/Brainsoft-Raxat/respire-api-go/config"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/data"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/models"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/repository"
	"github.com/Brainsoft-Raxat/respire-api-go/pkg/ctxconst"
	"go.uber.org/zap"
)

type challengeService struct {
	challengeRepo repository.ChallengeRepository
	sessionRepo   repository.SessionRepository
	userRepo      repository.UserRepository
	cfg           *config.Configs
	log           *zap.SugaredLogger
}

func NewChallengeService(repo *repository.Repository, cfg *config.Configs, log *zap.SugaredLogger) ChallengeService {
	return &challengeService{
		challengeRepo: repo.ChallengeRepository,
		sessionRepo:   repo.SessionRepository,
		userRepo:      repo.UserRepository,
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
		CreatedAt:       time.Now().UTC(),
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
	if req.ID == "" {
		req.ID = ctxconst.GetUserID(ctx)
	}

	challenge, err := s.challengeRepo.GetChallengeByID(ctx, req.ID)
	if err != nil {
		return data.GetChallengeByIDResponse{}, err
	}

	leaderboards := make([]*models.ShortUser, 0, len(challenge.Participants))

	for _, participant := range challenge.Participants {
		user, err := s.userRepo.GetUserByID(ctx, participant)
		if err != nil {
			continue
			// return data.GetChallengeByIDResponse{}, err
		}

		session, err := s.sessionRepo.GetSessionsByUserID(ctx, user.ID, challenge.CreatedAt, challenge.EndDate, 0)
		if err != nil {
			continue
			// return data.GetChallengeByIDResponse{}, err
		}

		leaderboards = append(leaderboards, &models.ShortUser{
			ID:         user.ID,
			Name:       user.Username,
			Username:   user.Username,
			Avatar:     user.Avatar,
			SmokeCount: session.Sum(),
		})

	}

	slices.SortFunc(leaderboards, func(a *models.ShortUser, b *models.ShortUser) int {
		if a.SmokeCount == b.SmokeCount {
			return strings.Compare(a.Username, b.Username)
		}
		return a.SmokeCount - b.SmokeCount
	})

	var prevSmokeCount int
	var prevPosition int
	for leaderboardsIdx, leaderboard := range leaderboards {
		if leaderboardsIdx == 0 {
			prevSmokeCount = leaderboard.SmokeCount
			prevPosition = leaderboardsIdx + 1
		} else {
			if leaderboard.SmokeCount == prevSmokeCount {
				leaderboard.Position = prevPosition
			} else {
				prevSmokeCount = leaderboard.SmokeCount
				prevPosition = leaderboardsIdx + 1
			}
		}
		leaderboard.Position = prevPosition
	}

	return data.GetChallengeByIDResponse{
		Challenge:   challenge,
		Leaderboard: leaderboards,
	}, nil
}

func (s *challengeService) GetChallengesByUserID(ctx context.Context, req data.GetChallengesByUserIDRequest) (data.GetChallengesByUserIDResponse, error) {
	if req.UserID == "" {
		req.UserID = ctxconst.GetUserID(ctx)
	}

	if req.From.IsZero() {
		req.From = time.Now().UTC()
	}

	challenges, err := s.challengeRepo.GetChallengesByUserID(ctx, req.UserID, req.ChallengeType, req.Invite, req.Limit, req.Page, req.From, req.To)
	if err != nil {
		return data.GetChallengesByUserIDResponse{}, err
	}

	for _, challenge := range challenges {
		owner, err := s.userRepo.GetUserByID(ctx, req.UserID)
		if err != nil {
			return data.GetChallengesByUserIDResponse{}, err
		}

		challenge.Owner = &models.ShortUser{
			ID:       owner.ID,
			Name:     owner.Name,
			Username: owner.Username,
			Avatar:   owner.Avatar,
		}

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
		UpdatedAt:       time.Now().UTC(),
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

func (s *challengeService) HandleChallengeInviation(ctx context.Context, req data.HandleChallengeInviationRequest) (data.HandleChallengeInviationResponse, error) {
	userID := ctxconst.GetUserID(ctx)

	challenge, err := s.challengeRepo.GetChallengeByID(ctx, req.ChallengeID)
	if err != nil {
		return data.HandleChallengeInviationResponse{}, err
	}

	if slices.Contains(challenge.Invited, userID) {
		if req.Accept {
			challenge.Participants = append(challenge.Participants, userID)
		}

		idx := slices.Index(challenge.Invited, userID)
		challenge.Invited = slices.Delete(challenge.Invited, idx, idx+1)
	}

	err = s.challengeRepo.UpdateChallenge(ctx, req.ChallengeID, challenge)
	if err != nil {
		return data.HandleChallengeInviationResponse{}, err
	}

	return data.HandleChallengeInviationResponse{
		ID: challenge.ID,
	}, nil
}
