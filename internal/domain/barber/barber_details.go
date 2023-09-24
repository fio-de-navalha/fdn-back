package barber

import "github.com/google/uuid"

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
		ID:       uuid.NewString(),
		BarberId: barberId,
		Address:  address,
	}
}

func NewContact(barberId string, contact string) *Contact {
	return &Contact{
		ID:       uuid.NewString(),
		BarberId: barberId,
		Contact:  contact,
	}
}
