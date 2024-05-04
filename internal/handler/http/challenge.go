package handler

import (
	"net/http"

	"github.com/Brainsoft-Raxat/respire-api-go/internal/data"
	"github.com/labstack/echo/v4"
)

// CreateChallenge godoc
// @Summary Create a challenge
// @Description Create a challenge
// @Tags challenges
// @Accept json
// @Produce json
// @Param challenge body data.CreateChallengeRequest true "Challenge object that needs to be created"
// @Success 200 {object} data.CreateChallengeResponse
// @Router /challenges [post]
func (h *handler) CreateChallenge(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	req := data.CreateChallengeRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	resp, err := h.service.ChallengeService.CreateChallenge(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

// GetChallengeByID godoc
// @Summary Get a challenge by ID
// @Description Get a challenge by ID
// @Tags challenges
// @Accept json
// @Produce json
// @Param id path string true "Challenge ID"
// @Success 200 {object} data.GetChallengeByIDResponse
// @Router /challenges/{id} [get]
func (h *handler) GetChallengeByID(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	req := data.GetChallengeByIDRequest{
		ID: c.Param("id"),
	}

	resp, err := h.service.ChallengeService.GetChallengeByID(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

// GetChallengesByUserID godoc
// @Summary Get challenges by user ID
// @Description Get challenges by user ID
// @Tags challenges
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} data.GetChallengesByUserIDResponse
// @Router /challenges/user/{id} [get]
func (h *handler) GetChallengesByUserID(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	req := data.GetChallengesByUserIDRequest{
		UserID:        c.Param("id"),
		Invite:        c.QueryParam("invite") == "true",
		// From:          time.Time{},
		// To:            time.Time{},
	}

	resp, err := h.service.ChallengeService.GetChallengesByUserID(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

// UpdateChallengeByID godoc
// @Summary Update a challenge by ID
// @Description Update a challenge by ID
// @Tags challenges
// @Accept json
// @Produce json
// @Param id path string true "Challenge ID"
// @Param challenge body data.UpdateChallengeRequest true "Challenge object that needs to be updated"
// @Success 200 {object} data.UpdateChallengeResponse
// @Router /challenges/{id} [put]
func (h *handler) UpdateChallengeByID(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	req := data.UpdateChallengeRequest{
		ID: c.Param("id"),
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	resp, err := h.service.ChallengeService.UpdateChallengeByID(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

// DeleteChallengeByID godoc
// @Summary Delete a challenge by ID
// @Description Delete a challenge by ID
// @Tags challenges
// @Accept json
// @Produce json
// @Param id path string true "Challenge ID"
// @Success 200
// @Router /challenges/{id} [delete]
func (h *handler) DeleteChallengeByID(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	req := data.GetChallengeByIDRequest{
		ID: c.Param("id"),
	}

	err := h.service.ChallengeService.DeleteChallengeByID(ctx, req)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}


// HandleChallengeInviation godoc
// @Summary Handle challenge inviation
// @Description Handle challenge inviation
// @Tags challenges
// @Accept json
// @Produce json
// @Param challenge body data.HandleChallengeInviationRequest true "Challenge object that needs to be handled"
// @Success 200 {object} data.HandleChallengeInviationResponse
// @Router /challenge/invitations/handle [post]
func (h *handler) HandleChallengeInviation(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	req := data.HandleChallengeInviationRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	resp, err := h.service.ChallengeService.HandleChallengeInviation(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}