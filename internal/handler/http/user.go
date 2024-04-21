package handler

import (
	"net/http"
	"strconv"

	"github.com/Brainsoft-Raxat/respire-api-go/internal/data"
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
func (h *handler) GetUserByID(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	req := data.GetUserByIDRequest{}

	req.ID = c.Param("id")

	user, err := h.service.UserService.GetUserByID(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
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
func (h *handler) GetUserByEmail(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	email := c.Param("email")

	user, err := h.service.UserService.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
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
func (h *handler) CreateUser(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	req := data.CreateUserRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp, err := h.service.UserService.CreateUser(ctx, req)
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
func (h *handler) UpdateUser(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	id := c.Param("id")

	req := data.UpdateUserRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp, err := h.service.UserService.UpdateUser(ctx, id, req)
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
func (h *handler) DeleteUser(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	id := c.Param("id")

	err := h.service.UserService.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusNoContent, nil)
}

// SearchUserByUsername godoc
// @Summary Search users by username
// @Description Search users by username
// @Tags users
// @Accept  json
// @Produce  json
// @Param username query string true "Username"
// @Param limit query int false "Limit"
// @Success 200 {object} data.SearchUsersByUsernameResponse
// @Failure 400 {string} string "Bad request"
// @Security BearerAuth
// @Router /user/search [get]
func (h *handler) SearchUserByUsername(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	req := data.SearchUsersByUsernameRequest{}

	req.Username = c.QueryParam("username")
	if limit := c.QueryParam("limit"); limit != "" {
		l, err := strconv.Atoi(limit)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		req.Limit = l
	}

	resp, err := h.service.UserService.SearchUsersByUsername(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}
