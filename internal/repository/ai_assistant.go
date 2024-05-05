package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Brainsoft-Raxat/respire-api-go/config"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/models"
	"github.com/Brainsoft-Raxat/respire-api-go/pkg/apperror"
	"github.com/Brainsoft-Raxat/respire-api-go/pkg/errcodes"
	"go.uber.org/zap"
)

type aiAssistantRepository struct {
	client *http.Client
	cfg    *config.AIAssistant
	logger *zap.SugaredLogger
}

func NewAIAssistantRepository(client *http.Client, cfg *config.AIAssistant, logger *zap.SugaredLogger) AIAssistantRepository {
	return &aiAssistantRepository{
		client: client,
		cfg:    cfg,
		logger: logger,
	}
}

func (r *aiAssistantRepository) GetRecommendations(ctx context.Context, req models.GetRecommendationRequest) (models.GetRecommendationResponse, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return models.GetRecommendationResponse{}, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to marshal request")
	}
	
	fmt.Println(r.cfg.BaseURL+"/api/v1/recommendations")

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, r.cfg.BaseURL+"/api/v1/recommendations", bytes.NewReader(reqBody))
	if err != nil {
		return models.GetRecommendationResponse{}, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to create request")
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := r.client.Do(httpReq)
	if err != nil {
		return models.GetRecommendationResponse{}, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to send request")
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.GetRecommendationResponse{}, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to read response body")
	}

	if resp.StatusCode != http.StatusOK {
		return models.GetRecommendationResponse{}, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to get recommendations")
	}

	var recommendationResp models.GetRecommendationResponse
	err = json.Unmarshal(respBody, &recommendationResp)
	if err != nil {
		return models.GetRecommendationResponse{}, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "failed to unmarshal response body")
	}

	return recommendationResp, nil
}
