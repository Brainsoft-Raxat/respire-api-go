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

type aiAssistantService struct {
	aiAssistantRepo repository.AIAssistantRepository
	cfg             *config.Configs
	log             *zap.SugaredLogger
}

func NewAIAssistantService(repo *repository.Repository, cfg *config.Configs, log *zap.SugaredLogger) AIAssistantService {
	return &aiAssistantService{
		aiAssistantRepo: repo.AIAssistantRepository,
		cfg:             cfg,
		log:             log,
	}
}

func (s *aiAssistantService) GetRecommendations(ctx context.Context, req data.GetRecommendationsRequest) (data.GetRecommendationsResponse, error) {
	resp, err := s.aiAssistantRepo.GetRecommendations(ctx, models.GetRecommendationRequest{
		EventType: "craving report",
		Data: models.GetRecommendationRequestData{
			CravingLevel: req.CravingLevel,
			Context:      req.Context,
			Mood:         req.Mood,
			Timestamp:    time.Now(),
		},
	})
	if err != nil {
		return data.GetRecommendationsResponse{}, err
	}

	return data.GetRecommendationsResponse{
		Reccomendations: resp.Recommendations,
	}, nil
}
