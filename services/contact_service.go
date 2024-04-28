package services

type ContactService struct {
	store *ContactStore
}

func NewContactService(store *ContactStore) *ContactService {
	return &ContactService{store: store}
}

func (s *ContactService) CreateContact(contact Contact) error {
	return s.store.CreateContact(contact)
}

func (s *ContactService) GetAllContacts() ([]Contact, error) {
	return s.store.GetAllContacts()
}

func (s *ContactService) GetContactById(id string) (Contact, error) {
	return s.store.GetContactById(id)
}

func (s *ContactService) UpdateContact(id string, contact Contact) error {
	return s.store.UpdateContact(id, contact)
}

func (s *ContactService) DeleteContact(id string) error {
	return s.store.DeleteContact(id)
}
