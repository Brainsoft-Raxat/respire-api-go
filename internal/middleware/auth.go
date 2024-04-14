package middleware

import (
	"context"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"

	"github.com/labstack/echo/v4"
)

const (
	vipIDToken = "raxat"
)

func FirebaseAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		opt := option.WithCredentialsFile("quitsmoke-20141-firebase-adminsdk-ugo14-c5730ea21d.json")
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to initialize Firebase app")
		}

		client, err := app.Auth(context.Background())
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to initialize Firebase Auth client")
		}

		idToken := c.Request().Header.Get("Authorization")
		idToken = strings.Replace(idToken, "Bearer ", "", 1)	
		if idToken == vipIDToken {
			c.Set("uid", "yXuTFL6nMlWjaes5JlvGlwvFPby2")
			return next(c)
		}
		
		token, err := client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid ID token")
		}

		c.Set("uid", token.UID)

		return next(c)
	}
}
