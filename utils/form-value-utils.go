package utils

import "github.com/iakovosvo/go-htmx-contacts-app/services"

func HasEmailError(formData services.FormData) bool {
	emailError, ok := formData.Errors["email"]

	if ok && emailError != "" {
		return true
	}
	return false
}

func HasFormValue(formData services.FormData, field string) bool {
	value, ok := formData.Values[field]

	if ok && len(value) > 0 && value != "" {
		return true
	}
	return false
}
