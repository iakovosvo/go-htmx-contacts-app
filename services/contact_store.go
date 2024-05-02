package services

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

var (
	ErrNotFound = errors.New("not found")
)

var ErrContactExists = errors.New("contact already exists")

type Contact struct {
	ID    string
	Name  string
	Email string
}

type ContactDataStore interface {
	CreateContact(contact Contact) error
	GetContactById(id string) (Contact, error)
	GetContacts(page, pageSize int) ([]Contact, error)
	UpdateContact(id string, contact Contact) error
	DeleteContact(id string) error
}

type ContactStore struct {
	Contacts []Contact
	Mutex    sync.Mutex
}

func NewContactsStore() *ContactStore {
	return &ContactStore{
		Contacts: []Contact{
			{
				Name:  "Mila Rivers",
				Email: "mila.rivers@example.com",
				ID:    uuid.New().String(),
			},
			{
				Name:  "Ethan Harper",
				Email: "ethan.harper@example.com",
				ID:    uuid.New().String(),
			},
			{
				Name:  "Isla Grey",
				Email: "isla.grey@example.com",
				ID:    uuid.New().String(),
			},
			{
				Name:  "Noah Lee",
				Email: "noah.lee@example.com",
				ID:    uuid.New().String(),
			},
			{
				Name:  "Olivia Smith",
				Email: "olivia.smith@example.com",
				ID:    uuid.New().String(),
			},
			{
				Name:  "Liam Johnson",
				Email: "liam.johnson@example.com",
				ID:    uuid.New().String(),
			},
			{
				Name:  "Emma Brown",
				Email: "emma.brown@example.com",
				ID:    uuid.New().String(),
			},
			{
				Name:  "Lucas Jones",
				Email: "lucas.jones@example.com",
				ID:    uuid.New().String(),
			},
			{
				Name:  "Sophia Garcia",
				Email: "sophia.garcia@example.com",
				ID:    uuid.New().String(),
			},
			{
				Name:  "Mason Martinez",
				Email: "mason.martinez@example.com",
				ID:    uuid.New().String(),
			},
		},
	}
}

func (c *ContactStore) CreateContact(contact Contact) error {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	for _, c := range c.Contacts {
		if c.Email == contact.Email {
			return ErrContactExists
		}
	}

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

func (c *ContactStore) GetContacts(page, pageSize int) ([]Contact, error) {
	start := (page - 1) * pageSize
	end := start + pageSize

	if end > len(c.Contacts) {
		end = len(c.Contacts)
	}

	if start >= len(c.Contacts) {
		return []Contact{}, nil
	}

	return c.Contacts[start:end], nil
}

func (c *ContactStore) UpdateContact(newContact Contact) error {
	for i, contact := range c.Contacts {
		if contact.ID == newContact.ID {
			c.Contacts[i] = newContact
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
