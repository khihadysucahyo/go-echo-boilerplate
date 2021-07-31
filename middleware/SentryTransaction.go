package middleware

import (
	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
)

type SentryMiddleware struct {
}

func (m *SentryMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Controrl-Allow-Origin", "*")

		return next(c)
	}
}

func (m *SentryMiddleware) SENTRY(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		span := sentry.StartSpan(
			c.Request().Context(), "", sentry.TransactionName(c.Request().URL.String()),
		)

		span.Finish()

		return next(c)
	}
}

func InitMiddleware() *SentryMiddleware {
	return &SentryMiddleware{}
}
