package handler

import (
	"net/http"

	"github.com/Brainsoft-Raxat/respire-api-go/internal/data"
	"github.com/labstack/echo/v4"
)

// GetRecommendations godoc
// @Summary Get recommendations
// @Description Get recommendations
// @Tags ai-assistant
// @Accept json
// @Produce json
// @Param recommendations body data.GetRecommendationsRequest true "Recommendations object that needs to be created"
// @Success 200 {object} data.GetRecommendationsResponse
// @Router /ai-assistant/recommendations [post]
func (h *handler) GetRecommendations(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	req := data.GetRecommendationsRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	resp, err := h.service.AIAssistantService.GetRecommendations(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}