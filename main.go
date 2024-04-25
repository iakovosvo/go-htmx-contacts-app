package main

import (
	"github.com/a-h/templ"
	"github.com/iakovosvo/go-htmx-todo-app/templates"
	"github.com/labstack/echo/v4"
)

func render(ctx echo.Context, cmp templ.Component) error {
	return cmp.Render(ctx.Request().Context(), ctx.Response())
}

func main() {
	e := echo.New()
	e.Static("/static", "static")

	e.GET("/", func(c echo.Context) error {
		return render(c, templates.HomePage())
	})
	e.Logger.Fatal(e.Start(":8080"))
}
