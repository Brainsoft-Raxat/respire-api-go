package handler

import (
	"net/http"

	"github.com/Brainsoft-Raxat/respire-api-go/config"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/middleware"
	"github.com/Brainsoft-Raxat/respire-api-go/internal/service"
	"github.com/labstack/echo/v4"
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
	e.Use(middleware.SetRequestContextWithTimeout(h.cfg.App.Timeout))
	api := e.Group("/api/v1")
	api.Use(middleware.FirebaseAuthMiddleware)
	{	
		user := api.Group("/user")
		{
			user.GET("/:id", h.GetUserByID)
			user.GET("/email/:email", h.GetUserByEmail)
			user.POST("", h.CreateUser)
			user.PUT("/:id", h.UpdateUser)
			user.DELETE("/:id", h.DeleteUser)
		}
		api.GET("", h.handleExample)
	}
}

func (h *handler) handleExample(c echo.Context) error {
	// ctx := c.Request().Context()
	return c.String(http.StatusOK, "Hello, World!")
}
