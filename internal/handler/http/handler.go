package handler

import (
	"context"
	"net/http"

	"github.com/Brainsoft-Raxat/respire-api-go/config"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/middleware"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/service"
	"github.com/Brainsoft-Raxat/respire-api-go/pkg/apperror"
	"github.com/Brainsoft-Raxat/respire-api-go/pkg/ctxconst"
	"github.com/Brainsoft-Raxat/respire-api-go/pkg/errcodes"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	"go.uber.org/zap"
)

type handler struct {
	service *service.Service
	cfg     *config.Configs
	logger  *zap.SugaredLogger
}

type Handler interface {
	SetAPI(e *echo.Echo)
}

func New(services *service.Service, cfg *config.Configs, logger *zap.SugaredLogger) Handler {
	return &handler{
		service: services,
		cfg:     cfg,
		logger:  logger,
	}
}

func (h *handler) SetAPI(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Use(middleware.RequestIDMiddleware)
	e.Use(middleware.BodyDumpMiddleware)
	e.Use(errorHandler(h.logger))
	e.Use(middleware.CustomLogger(h.logger))
	api := e.Group("/api/v1")
	api.Use(middleware.FirebaseAuthMiddleware)
	{
		user := api.Group("/user")
		{
			user.POST("", h.CreateUser)
			user.GET("/search", h.SearchUserByUsername)
			user.GET("/:id", h.GetUserByID)
			user.PUT("/:id", h.UpdateUser)
			user.DELETE("/:id", h.DeleteUser)
			user.GET("/by-email/:email", h.GetUserByEmail)
		}

		friends := api.Group("/friends")
		{
			friends.POST("/invitations", h.CreateFriendshipInvitation)
			friends.GET("/invitations", h.GetFriendshipInvitations)
			friends.POST("/invitations/handle", h.HandleFriendshipInvitation)
			friends.GET("", h.GetFriendsList)
			friends.DELETE("/:id", h.RemoveFriend)
		}

		api.GET("", h.handleExample)
	}
}

func (h *handler) handleExample(c echo.Context) error {
	// ctx := c.Request().Context()
	return c.String(http.StatusOK, "Hello, World!")
}

func errorHandler(logger *zap.SugaredLogger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err != nil {
				if appErr := apperror.AsErrorInfo(err); appErr != nil {

					errorDetails := map[string]interface{}{
						"code":             appErr.Code,
						"status":           appErr.Status,
						"message":          appErr.Message,
						"developerMessage": appErr.DeveloperMessage,
					}

					logger.Errorw("response error",
						"request_id", c.Get("request_id"),
						"status", appErr.Status,
						"error", errorDetails,
					)

					appErr.DeveloperMessage = ""

					return c.JSON(appErr.Status, appErr)
				}

				logger.Errorw("unexpected error",
					"request_id", c.Get("request_id"),
					"error", err,
				)

				return c.JSON(http.StatusInternalServerError, apperror.NewErrorInfo(c.Request().Context(), errcodes.InternalServerError, ""))
			}

			return nil
		}
	}
}

func (h *handler) context(c echo.Context) (context.Context, context.CancelFunc) {
	ctx := c.Request().Context()

	ctx = ctxconst.SetUserID(ctx, c.Get("uid").(string))

	return context.WithTimeout(ctx, h.cfg.App.Timeout)
}
