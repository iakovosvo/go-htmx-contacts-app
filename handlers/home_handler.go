package handlers

import (
	"github.com/iakovosvo/go-htmx-contacts-app/services"
)

// Assuming ContactService is defined in the 'services' package

type HomeHandler struct {
	service services.Contact
}

func NewHomeHandler(service services.Contact) *HomeHandler {
	return &HomeHandler{
		service: service,
	}
}

// func (h *HomeHandler) ListContacts(c echo.Context) error {
//     contacts, err := h.service.
//     // if err != nil {
//     //     return c.JSON(http.StatusInternalServerError, map[string]string{
//     //         "error": "Failed to retrieve contacts",
//     //     })
//     // }
//    // return c.JSON(http.StatusOK, contacts)
// }
