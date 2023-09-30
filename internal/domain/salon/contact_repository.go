package salon

type ContactRepository interface {
	FindBySalonId(salonId string) ([]*Contact, error)
	FindById(id string, salonId string) (*Contact, error)
	Save(contact *Contact) (*Contact, error)
	Delete(contactId string) error
}
