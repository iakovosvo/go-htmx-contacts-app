package middleware

import "github.com/labstack/echo/v4"

func SetContentType(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
		return next(c)
	}
}
