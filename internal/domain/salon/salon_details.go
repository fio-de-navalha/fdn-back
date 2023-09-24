package salon

import "github.com/google/uuid"

type Address struct {
	ID      string `json:"id"`
	SalonId string `json:"salon_id"`
	Address string `json:"address"`
}

type Contact struct {
	ID      string `json:"id"`
	SalonId string `json:"salon_id"`
	Contact string `json:"contact"`
}

func NewAddress(salonId string, address string) *Address {
	return &Address{
		ID:      uuid.NewString(),
		SalonId: salonId,
		Address: address,
	}
}

func NewContact(salonId string, contact string) *Contact {
	return &Contact{
		ID:      uuid.NewString(),
		SalonId: salonId,
		Contact: contact,
	}
}
