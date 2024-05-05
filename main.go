package main

import (
	"github.com/iakovosvo/go-htmx-contacts-app/handlers"
	"github.com/iakovosvo/go-htmx-contacts-app/middleware"
	"github.com/iakovosvo/go-htmx-contacts-app/services"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	store := services.NewContactsStore()
	service := services.NewContactService(store)
	contactsHandler := handlers.NewContactsHandler(service)

	// TODO: url & redirection

	// TODO: Take a look at The Primagen 422 script

	// TODO: Add placeholder text if list is empty

	e.Static("/static", "static")
	e.Use(middleware.SetContentType)

	e.GET("/", contactsHandler.HomeHandler)
	e.GET("/contacts", contactsHandler.GetContacts)
	e.GET("/contacts/:id", contactsHandler.GetContact)
	e.POST("/contacts", contactsHandler.Create)
	e.POST("/contacts/:id", contactsHandler.Update)
	e.DELETE("/contacts/:id", contactsHandler.Delete)
	e.Logger.Fatal(e.Start(":8080"))
}
