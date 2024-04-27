package services

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

var (
	ErrNotFound = errors.New("not found")
)

// TODO: Move to models
type Contact struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
}

type ContactStore struct {
	Contacts []Contact
	Mutex    sync.Mutex
}

func NewContactsStore() *ContactStore {
	return &ContactStore{
		Contacts: []Contact{
			{FirstName: "John", LastName: "Johnson", Email: "jj@gmail.com", ID: uuid.New().String()},
		},
	}
}

func (c *ContactStore) CreateContact(contact Contact) error {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.Contacts = append(c.Contacts, contact)
	return nil
}

func (c *ContactStore) GetContactById(id string) (Contact, error) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	for _, contact := range c.Contacts {
		if contact.ID == id {
			return contact, nil
		}
	}
	return Contact{}, ErrNotFound
}

func (c *ContactStore) GetAllContacts() ([]Contact, error) {
	return c.Contacts, nil
}

func (c *ContactStore) UpdateContact(id string, contact Contact) error {
	for i, contact := range c.Contacts {
		if contact.ID == id {
			c.Contacts[i].ID = id
			return nil
		}
	}
	return ErrNotFound
}

func (c *ContactStore) DeleteContact(id string) error {
	for i, contact := range c.Contacts {
		if contact.ID == id {
			c.Contacts = append(c.Contacts[:i], c.Contacts[i+1:]...)
			return nil
		}
	}
	return nil
}

// func (c *ContactStore) ContactExists(email string) bool {
// 	c.Mutex.Lock()
// 	defer c.Mutex.Unlock()
// 	for _, contact := range c.Contacts {
// 		if contact.Email == email {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (app *ContactStore) InjectData(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		c.Set("data", app)
// 		return next(c)
// 	}
// }
