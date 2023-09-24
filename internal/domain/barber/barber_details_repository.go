package barber

type AddressRepository interface {
	FindByBarberId(barberId string) ([]*Address, error)
	FindById(id string) (*Address, error)
	Save(address *Address) (*Address, error)
	Delete(addressId string) error
}

type ContactRepository interface {
	FindByBarberId(barberId string) ([]*Contact, error)
	FindById(id string) (*Contact, error)
	Save(contact *Contact) (*Contact, error)
	Delete(contactId string) error
}
