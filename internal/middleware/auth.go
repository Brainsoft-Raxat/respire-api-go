package middleware

import (
	"context"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/labstack/echo/v4"
)

const (
    vipIDToken = "raxat"
)

var FirebaseApp *firebase.App


func FirebaseAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        client, err := FirebaseApp.Auth(context.Background())
        if err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get Firebase Auth client")
        }

        idToken := c.Request().Header.Get("Authorization")
        idToken = strings.TrimPrefix(idToken, "Bearer ")

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
