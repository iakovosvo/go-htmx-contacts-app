package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/iakovosvo/go-htmx-contacts-app/services"
	"github.com/iakovosvo/go-htmx-contacts-app/templates"
	"github.com/labstack/echo/v4"
)

type ContactService interface {
	CreateContact(c services.Contact) error
	// GetContactById(c services.Contact) error
	GetAllContacts() ([]services.Contact, error)
	// UpdateContact(c services.Contact) error
	// DeleteContact(c services.Contact) error
}

type ContactsHandler struct {
	store ContactService
}

func NewContactService(s ContactService) *ContactsHandler {
	return &ContactsHandler{
		store: s,
	}
}

func (h *ContactsHandler) Create(c echo.Context) error {
	firstName := c.FormValue("firstName")
	lastName := c.FormValue("lastName")
	email := c.FormValue("email")

	//handler := c.Get("data").(*ContactsHandler)

	// if ContactExists(handler, email) {
	// 	formData := services.FormData{
	// 		Errors: map[string]string{
	// 			"email": "Email already exists",
	// 		},
	// 		Values: map[string]string{
	// 			"firstName": firstName,
	// 			"lastName":  lastName,
	// 			"email":     email,
	// 		},
	// 	}
	// 	c.Response().Writer.WriteHeader(http.StatusUnprocessableEntity)
	// 	if err := Render(c, templates.Form(formData)); err != nil {
	// 		// Log the error and handle it appropriately
	// 		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to render form with errors")
	// 	}
	// 	return nil
	// }
	// TODO: Make a NewContact handler
	contact := services.Contact{ID: uuid.New().String(), FirstName: firstName, LastName: lastName, Email: email}
	if err := h.store.CreateContact(contact); err != nil {
		return err
	}

	formData := services.NewFormData()
	if err := Render(c, templates.Form(formData)); err != nil {
		// Log the error and handle it appropriately
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to render form")
	}
	return nil
}

func (h *ContactsHandler) GetAll(c echo.Context) error {
	contacts, err := h.store.GetAllContacts()

	if err != nil {
		return err
	}

	if err := Render(c, templates.ContactList(contacts)); err != nil {
		// Log the error and handle it appropriately
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to render list")
	}
	return nil
}

func (h *ContactsHandler) HomeHandler(c echo.Context) error {
	contacts, err := h.store.GetAllContacts()
	formData := services.NewFormData()

	if err != nil {
		return err
	}

	if err := Render(c, templates.HomePage(contacts, formData)); err != nil {
		// Log the error and handle it appropriately
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to render list")
	}
	return nil
}
