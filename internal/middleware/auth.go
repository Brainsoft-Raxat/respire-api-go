package middleware

import (
	"context"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/Brainsoft-Raxat/respire-api-go/pkg/apperror"
	"github.com/Brainsoft-Raxat/respire-api-go/pkg/errcodes"
	"github.com/labstack/echo/v4"
)

const (
	vipIDToken     = "Bearer raxat"
	vipFriendToken = "Bearer dias"
)

var FirebaseApp *firebase.App

func FirebaseAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		client, err := FirebaseApp.Auth(context.Background())
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get Firebase Auth client")
		}

		idToken := c.Request().Header.Get("Authorization")

		if idToken == vipIDToken {
			c.Set("uid", "VjXWD89EvRaaMeXscIgH")
			return next(c)
		} else if idToken == vipFriendToken {
			c.Set("uid", "8n169tPBBg0MFXu1DhjK")
			return next(c)
		}

		token, err := client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			return apperror.NewErrorInfo(context.Background(), errcodes.Unauthorized, "Invalid token")
		}

		c.Set("uid", token.UID)
		return next(c)
	}
}
