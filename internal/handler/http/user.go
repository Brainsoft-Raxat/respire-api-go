package handler

import (
	"net/http"

	"github.com/Brainsoft-Raxat/respire-api-go/internal/data"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/models"
	"github.com/labstack/echo/v4"
)

func (h *handler) GetUserByID(c echo.Context) error {
	id := c.Param("id")

	user, err := h.service.UserService.GetUserByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h *handler) GetUserByEmail(c echo.Context) error {
	email := c.Param("email")

	user, err := h.service.UserService.GetUserByEmail(c.Request().Context(), email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h *handler) CreateUser(c echo.Context) error {
	req := data.CreateUserRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp, err := h.service.UserService.CreateUser(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, resp)
}

func (h *handler) UpdateUser(c echo.Context) error {
	id := c.Param("id")

	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := h.service.UserService.UpdateUser(c.Request().Context(), id, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h *handler) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	err := h.service.UserService.DeleteUser(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusNoContent, nil)
}
