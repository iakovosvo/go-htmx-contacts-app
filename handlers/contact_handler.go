package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/google/uuid"
	"github.com/iakovosvo/go-htmx-contacts-app/services"
	"github.com/iakovosvo/go-htmx-contacts-app/views/components"
	"github.com/iakovosvo/go-htmx-contacts-app/views/pages"

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
			return Render(c, components.Form(formData))
		}
	}

	formData := services.NewFormData()
	c.Response().Header().Set("Content-Type", "text/html")

	if err := Render(c, components.Form(formData)); err != nil {
		// Log the error and handle it appropriately
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to render form")
	}

	return Render(c, components.ContactItemOob(contact))
}

func (h *ContactsHandler) Update(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	id := c.Param("id")

	contact := services.Contact{ID: id, Name: name, Email: email}

	// TODO: Handle errors
	h.service.UpdateContact(contact)

	formData := services.NewFormData()
	c.Response().Header().Set("Content-Type", "text/html")

	if err := Render(c, components.Form(formData)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to render form")
	}

	return Render(c, components.UpdateContactItemOob(contact))
}

func (h *ContactsHandler) Delete(c echo.Context) error {
	id := c.Param("id")

	h.service.DeleteContact(id)
	// TODO: Add error handling

	return c.NoContent(200)
}

func (h *ContactsHandler) GetContact(c echo.Context) error {
	id := c.Param("id")

	contact, err := h.service.GetContactById(id)

	if err != nil {
		// TODO: handle errors and status code
		return err
	}

	formData := services.FormData{
		Values: map[string]string{
			"name":  contact.Name,
			"email": contact.Email,
			"id":    contact.ID,
		},
	}
	// TODO: For demo show the JSON response
	c.Response().Header().Set("Content-Type", "text/html")
	return Render(c, components.Form(formData))
}

func (h *ContactsHandler) GetContacts(c echo.Context) error {
	pageParam := c.QueryParam("page")
	pageSizeParam := 3

	page, err := strconv.Atoi(pageParam)

	if err != nil || page < 1 {
		page = 1
	}

	contacts, err := h.service.GetContacts(page, pageSizeParam)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	c.Response().Header().Set("Content-Type", "text/html")

	for i, contact := range contacts {
		attrs := templ.Attributes{"": ""}

		if i == len(contacts)-1 {
			attrs = templ.Attributes{
				"hx-get":     fmt.Sprintf("/contacts?page=%d", page+1),
				"hx-trigger": "revealed",
				"hx-target":  "#contact-list",
				"hx-swap":    "beforeend swap:500ms",
			}
		}
		Render(c, components.ContactItem(contact, attrs))
	}

	return nil
}

func (h *ContactsHandler) HomeHandler(c echo.Context) error {
	formData := services.NewFormData()

	if err := Render(c, pages.HomePage(formData)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to render home page")
	}

	return nil
}
