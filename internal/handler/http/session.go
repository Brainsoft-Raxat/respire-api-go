package handler

import (
	"net/http"

	"github.com/Brainsoft-Raxat/respire-api-go/internal/data"
	"github.com/labstack/echo/v4"
)

// GetSessionByID godoc
// @Summary Get a session by ID
// @Description Get a session by ID
// @Tags sessions
// @Accept json
// @Produce json
// @Param id path string true "Session ID"
// @Success 200 {object} data.GetSessionByIDResponse
// @Router /sessions/{id} [get]
func (h *handler) GetSessionByID(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	req := data.GetSessionByIDRequest{}

	req.ID = c.Param("id")

	session, err := h.service.SessionService.GetSessionByID(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, session)
}

// GetSessionByUserID godoc
// @Summary Get a session by user ID
// @Description Get a session by user ID
// @Tags sessions
// @Accept json
// @Produce json
// @Success 200 {object} data.GetSessionByUserIDResponse
// @Router /sessions/{uid} [get]
func (h *handler) GetSessionByUserID(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	uid := c.Param("uid")

	session, err := h.service.SessionService.GetSessionByUserID(ctx, data.GetSessionByUserIDRequest{ID: uid})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, session)
}

// GetSessionByTime godoc
// @Summary Get a session by user ID and time
// @Description Get a session by user ID and time
// @Tags sessions
// @Accept json
// @Produce json
// @Param time query data.GetSessionByUserIDAndDateRequest true "Time Query"
// @Success 200 {object} data.GetSessionByUserIDAndDateResponse
// @Router /sessions/by_time/ [get]
func (h *handler) GetSessionByTime(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	req := data.GetSessionByUserIDAndDateRequest{
		ID: c.Param("uid"),
	}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	resp, err := h.service.SessionService.GetSessionsByUserIDAndDateRange(ctx, req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp.Sum)
}

// CreateSession godoc
// @Summary Create a session
// @Description Create a session
// @Tags sessions
// @Accept json
// @Produce json
// @Param session body data.CreateSessionRequest true "Session object that needs to be created"
// @Success 200 {object} data.CreateSessionResponse
// @Router /sessions [post]
func (h *handler) CreateSession(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	req := data.CreateSessionRequest{UID: c.Param("uid")}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp, err := h.service.SessionService.CreateSession(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, resp)
}

// UpdateSession godoc
// @Summary Update a session
// @Description Update a session
// @Tags sessions
// @Accept json
// @Produce json
// @Param id path string true "Session ID"
// @Param session body data.UpdateSessionRequest true "Session object that needs to be updated"
// @Success 200 {object} data.UpdateSessionResponse
// @Router /sessions/{id} [put]
func (h *handler) UpdateSession(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	id := c.Param("id")

	req := data.UpdateSessionRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp, err := h.service.SessionService.UpdateSession(ctx, id, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

// DeleteSession godoc
// @Summary Delete a session
// @Description Delete a session
// @Tags sessions
// @Accept json
// @Produce json
// @Param id path string true "Session ID"
// @Success 204
// @Router /sessions/{id} [delete]
func (h *handler) DeleteSession(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	id := c.Param("id")

	err := h.service.SessionService.DeleteSession(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusNoContent, nil)
}

// GetUserStat godoc
// @Summary Get User statistics
// @Description Get User statistics
// @Tags sessions
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 204
// @Router /sessions/stat/{id} [get]
func (h *handler) GetUserStat(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()
	req := data.GetUserStatRequest{}
	req.ID = c.Param("id")

	stats, err := h.service.SessionService.GetUserStat(ctx, req)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, stats)
}
