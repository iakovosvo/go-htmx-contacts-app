package handlers

import (
	"github.com/iakovosvo/go-htmx-contacts-app/services"
)

type HomeHandler struct {
	service services.Contact
}

func NewHomeHandler(service services.Contact) *HomeHandler {
	return &HomeHandler{
		service: service,
	}
}
