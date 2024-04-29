package main

import (
	"github.com/iakovosvo/go-htmx-contacts-app/handlers"
	"github.com/iakovosvo/go-htmx-contacts-app/services"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	store := services.NewContactsStore()
	service := services.NewContactService(store)
	contactsHandler := handlers.NewContactsHandler(service)

	// TODO: hx-boost, url & redirection

	// TODO: Take a look at The Primagen 422 script

	e.Static("/static", "static")

	e.GET("/", contactsHandler.HomeHandler)
	e.GET("/contacts", contactsHandler.GetAll)
	e.GET("/contacts/:id", contactsHandler.Get)
	e.POST("/contacts", contactsHandler.Create)
	e.DELETE("/contacts/:id", contactsHandler.Delete)
	e.Logger.Fatal(e.Start(":8080"))
}
