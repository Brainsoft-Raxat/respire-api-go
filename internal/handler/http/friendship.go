package handler

import (
	"net/http"
	"strconv"

	"github.com/Brainsoft-Raxat/respire-api-go/internal/data"
	"github.com/labstack/echo/v4"
)

// CreateFriendshipInvitation godoc
// @Summary Create friendship invitation
// @Description Create friendship invitation
// @Tags friends
// @Accept  json
// @Produce  json
// @Param request body data.CreateFriendshipInvitationRequest true "Create Friendship Invitation Request"
// @Success 200 {object} data.CreateFriendshipInvitationResponse
// @Failure 400 {string} string "Invalid request"
// @Security BearerAuth
// @Router /friends/invitations [post]
func (h *handler) CreateFriendshipInvitation(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	req := data.CreateFriendshipInvitationRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	resp, err := h.service.FriendshipService.CreateFriendshipInvitation(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

// GetFriendshipInvitations godoc
// @Summary Get friendship invitations
// @Description Get friendship invitations
// @Tags friends
// @Accept  json
// @Produce  json
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {object} data.GetFriendshipInvitationsResponse
// @Failure 400 {string} string "Invalid limit or page"
// @Security BearerAuth
// @Router /friends/invitations [get]
func (h *handler) GetFriendshipInvitations(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	req := data.GetFriendshipInvitationsRequest{}

	limit := c.QueryParam("limit")
	if limit != "" {
		l, err := strconv.Atoi(limit)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid limit")
		}
		req.Limit = l
	}

	page := c.QueryParam("page")
	if page != "" {
		p, err := strconv.Atoi(page)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid page")
		}
		req.Page = p
	}

	resp, err := h.service.FriendshipService.GetFriendshipInvitations(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

// HandleFriendshipInvitation godoc
// @Summary Handle friendship invitation
// @Description Handle friendship invitation
// @Tags friends
// @Accept  json
// @Produce  json
// @Param request body data.HandleFriendshipInvitationRequest true "Handle Friendship Invitation Request"
// @Success 200 {object} data.HandleFriendshipInvitationResponse
// @Failure 400 {string} string "Invalid request"
// @Security BearerAuth
// @Router /friends/invitations/handle [post]
func (h *handler) HandleFriendshipInvitation(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	req := data.HandleFriendshipInvitationRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	resp, err := h.service.FriendshipService.HandleFriendshipInvitation(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

// GetFriendsList godoc
// @Summary Get friends list
// @Description Get friends list by user ID
// @Tags friends
// @Accept  json
// @Produce  json
// @Param user_id query string false "User ID (optional)"
// @Success 200 {object} data.GetFriendsListResponse
// @Failure 400 {string} string "Bad request"
// @Security BearerAuth
// @Router /friends [get]
func (h *handler) GetFriendsList(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	req := data.GetFriendsListRequest{}

	req.UserID = c.QueryParam("user_id")

	resp, err := h.service.FriendshipService.GetFriendsList(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

// RemoveFriend godoc
// @Summary Remove friend
// @Description Remove friend
// @Tags friends
// @Accept  json
// @Produce  json
// @Param id path string true "Friend ID"
// @Success 200 {object} data.RemoveFriendResponse
// @Failure 400 {string} string "Invalid request"
// @Security BearerAuth
// @Router /friends/{id} [delete]
func (h *handler) RemoveFriend(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	req := data.RemoveFriendRequest{
		FriendID: c.Param("id"),
	}

	resp, err := h.service.FriendshipService.RemoveFriend(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}
