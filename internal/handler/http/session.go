package handler

import (
	"net/http"
	"time"

	"github.com/Brainsoft-Raxat/respire-api-go/internal/data"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/repository"
	"github.com/labstack/echo/v4"
)

// GetUserByID godoc
// @Summary Get user by ID
// @Description Retrieve a user by their ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} data.GetUserByIDResponse
// @Failure 404 {string} string "User not found"
// @Security BearerAuth
// @Router /user/{id} [get]
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

// GetUserByEmail godoc
// @Summary Get user by email
// @Description Retrieve a user by their email
// @Tags users
// @Accept  json
// @Produce  json
// @Param email path string true "User email"
// @Success 200 {object} models.User
// @Failure 404 {string} string "User not found"
// @Security BearerAuth
// @Router /user/by-email/{email} [get]
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

func (h *handler) GetSessionByTime(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	uid := c.Param("uid")
	ti := c.Param("time")
	timeRange := [2]time.Time{}
	now := time.Now()
	switch ti {
	case "week":
		timeRange = repository.GetWeek(now)
	case "month":
		timeRange = repository.GetMonth(now)

	}

	session, err := h.service.SessionService.GetSessionsByUserIDAndDateRange(ctx, data.GetSessionByUserIDAndDateRequest{ID: uid, DR: timeRange})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, session)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body data.CreateUserRequest true "User object"
// @Success 201 {object} data.CreateUserResponse
// @Failure 400 {string} string "Bad request"
// @Security BearerAuth
// @Router /user [post]
func (h *handler) CreateSession(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	req := data.CreateSessionRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp, err := h.service.SessionService.CreateSession(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, resp)
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update a user by their ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param user body data.UpdateUserRequest true "User object"
// @Success 200 {object} data.UpdateUserResponse
// @Failure 400 {string} string "Bad request"
// @Security BearerAuth
// @Router /user/{id} [put]
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

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user by their ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {string} string "Bad request"
// @Security BearerAuth
// @Router /user/{id} [delete]
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

// SearchUserByUsername godoc
// @Summary Search users by username
// @Description Search users by username
// // @Tags users
// // @Accept  json
// // @Produce  json
// // @Param username query string true "Username"
// // @Param limit query int false "Limit"
// // @Success 200 {object} data.SearchUsersByUsernameResponse
// // @Failure 400 {string} string "Bad request"
// // @Security BearerAuth
// // @Router /user/search [get]
// func (h *handler) SearchUserByUsername(c echo.Context) error {
// 	ctx, cancel := h.context(c)
// 	defer cancel()

// 	req := data.SearchUsersByUsernameRequest{}

// 	req.Username = c.QueryParam("username")
// 	if limit := c.QueryParam("limit"); limit != "" {
// 		l, err := strconv.Atoi(limit)
// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, err.Error())
// 		}
// 		req.Limit = l
// 	}

// 	resp, err := h.service.UserService.SearchUsersByUsername(ctx, req)
// 	if err != nil {
// 		return err
// 	}

// 	return c.JSON(http.StatusOK, resp)
// }
