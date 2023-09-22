package barber

type Address struct {
	ID       string `json:"id"`
	BarberId string `json:"barber_id"`
	Address  string `json:"address"`
}

type Contact struct {
	ID       string `json:"id"`
	BarberId string `json:"barber_id"`
	Contact  string `json:"contact"`
}

func NewAddress(barberId string, address string) *Address {
	return &Address{
		BarberId: barberId,
		Address:  address,
	}
}

func NewContact(barberId string, contact string) *Contact {
	return &Contact{
		BarberId: barberId,
		Contact:  contact,
	}
}

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
