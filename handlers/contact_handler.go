package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/iakovosvo/go-htmx-contacts-app/services"
	"github.com/iakovosvo/go-htmx-contacts-app/templates"
	"github.com/labstack/echo/v4"
)

type ContactsHandler struct {
	service *services.ContactService
}

func NewContactsHandler(service *services.ContactService) *ContactsHandler {
	return &ContactsHandler{service: service}
}

func (h *ContactsHandler) Create(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")

	contact := services.Contact{ID: uuid.New().String(), Name: name, Email: email}

	if err := h.service.CreateContact(contact); err != nil {
		if err == services.ErrContactExists {
			formData := services.FormData{
				Errors: map[string]string{
					"email": "Email already exists",
				},
				Values: map[string]string{
					"name":  name,
					"email": email,
				},
			}
			// TODO: Handle http header
			return Render(c, templates.Form(formData))
		}
	}

	formData := services.NewFormData()
	c.Response().Header().Set("Content-Type", "text/html")

	if err := Render(c, templates.Form(formData)); err != nil {
		// Log the error and handle it appropriately
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to render form")
	}

	return Render(c, templates.ContactItemOob(contact))
}

func (h *ContactsHandler) Delete(c echo.Context) error {
	id := c.Param("id")

	h.service.DeleteContact(id)
	// TODO: Add error handling

	return c.NoContent(200)
}

func (h *ContactsHandler) GetAll(c echo.Context) error {
	contacts, err := h.service.GetAllContacts()

	if err != nil {
		return err
	}

	// TODO: For demo show the JSON response
	if err := Render(c, templates.ContactList(contacts)); err != nil {
		// Log the error and handle it appropriately
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to render list")
	}
	return nil
}

func (h *ContactsHandler) HomeHandler(c echo.Context) error {
	contacts, err := h.service.GetAllContacts()
	formData := services.NewFormData()

	if err != nil {
		return err
	}

	if err := Render(c, templates.HomePage(contacts, formData)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to render list")
	}
	return nil
}
