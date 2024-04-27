package services

type FormData struct {
	Errors map[string]string
	Values map[string]string
}

func NewFormData() FormData {
	return FormData{
		Errors: map[string]string{},
		Values: map[string]string{},
	}
}

// type PageData struct {
// 	Contacts ContactStore
// 	Form     FormData
// }

// func NewPageData(contacts ContactStore, form FormData) PageData {
// 	return PageData{
// 		Contacts: contacts,
// 		Form:     form,
// 	}
// }
