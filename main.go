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

	e.Static("/static", "static")
	r := e.Group("")
	r.Use(middleware.SetContentType)

	r.GET("/", contactsHandler.HomeHandler)
	r.GET("/contacts", contactsHandler.GetContacts)
	r.GET("/contacts/:id", contactsHandler.GetContact)
	r.POST("/contacts", contactsHandler.Create)
	r.POST("/contacts/:id", contactsHandler.Update)
	r.DELETE("/contacts/:id", contactsHandler.Delete)
	e.Logger.Fatal(e.Start(":8080"))
}
