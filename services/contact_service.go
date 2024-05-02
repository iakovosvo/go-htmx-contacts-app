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

func (s *ContactService) GetContacts(page, pageSize int) ([]Contact, error) {
	return s.store.GetContacts(page, pageSize)
}

func (s *ContactService) GetContactById(id string) (Contact, error) {
	return s.store.GetContactById(id)
}

func (s *ContactService) UpdateContact(contact Contact) error {
	return s.store.UpdateContact(contact)
}

func (s *ContactService) DeleteContact(id string) error {
	return s.store.DeleteContact(id)
}
