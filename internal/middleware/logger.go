package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func RequestIDMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		reqID := uuid.New().String()
		c.Set("request_id", reqID)
		c.Request().Header.Set("X-Request-ID", reqID)
		return next(c)
	}
}

func BodyDumpMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		if req.Body != nil {
			bodyBytes, err := io.ReadAll(req.Body)
			if err != nil {
				return err
			}

			req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			var bodyObj map[string]interface{}
			if err := json.Unmarshal(bodyBytes, &bodyObj); err != nil {
				c.Set("request_body", string(bodyBytes))
			} else {
				c.Set("request_body", bodyObj)
			}
		}

		return next(c)
	}
}

type responseBodyWriter struct {
	http.ResponseWriter
	body *bytes.Buffer
}

func (w *responseBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func CustomLogger(logger *zap.SugaredLogger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			originalWriter := c.Response().Writer
			responseBody := new(bytes.Buffer)

			c.Response().Writer = &responseBodyWriter{
				ResponseWriter: originalWriter,
				body:           responseBody,
			}

			requestBody := c.Get("request_body")

			queries := make(map[string][]string)
			for k, v := range c.QueryParams() {
				queries[k] = v
			}

			start := time.Now()
			err := next(c)
			duration := time.Since(start)

			var responseBodyObject interface{}
			if json.Unmarshal(responseBody.Bytes(), &responseBodyObject) != nil {
				responseBodyObject = responseBody.String()
			}

			logger.Infow("request",
				"request_id", c.Get("request_id"),
				"method", req.Method,
				"path", req.URL.Path,
				"body", requestBody,
				"queries", queries,
			)

			if err != nil {
				return err
			}

			logger.Infow("response",
				"request_id", c.Get("request_id"),
				"status", c.Response().Status,
				"duration", duration,
				"body", responseBodyObject,
			)

			c.Response().Writer = originalWriter

			return err
		}
	}
}
