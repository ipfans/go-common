package echoext

import (
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// ExtStdContext 扩展标准context.Context信息，方便后续Logger使用
func ExtStdContext(fn func(c echo.Context) interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			reqID := c.Request().Header.Get("X-Request-ID")
			if reqID == "" {
				reqID = uuid.New().String()
			}
			var id interface{}
			if fn != nil {
				id = fn(c)
			}
			ctx = context.WithValue(ctx, "logger_client_ip", c.RealIP())
			ctx = context.WithValue(ctx, "logger_request_id", reqID)
			ctx = context.WithValue(ctx, "logger_unique_id", id)
			req := c.Request().WithContext(ctx)
			c.SetRequest(req)
			err := next(c)
			c.Response().Header().Set("X-Request-ID", reqID)
			return err
		}
	}
}
