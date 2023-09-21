package barber

type ContactRepository interface {
	FindMany() ([]*Contact, error)
	FindById(id string) (*Contact, error)
	Save(contact *Contact) (*Contact, error)
}
