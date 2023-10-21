package salon

import (
	"github.com/google/uuid"
)

type AddSalonAddressRequest struct {
	Address string `json:"address"`
}

type Address struct {
	ID      string `json:"id"`
	SalonId string `json:"salonId"`
	Address string `json:"address"`

	Salon *Salon
}

func NewAddress(salonId string, address string) *Address {
	return &Address{
		ID:      uuid.NewString(),
		SalonId: salonId,
		Address: address,
	}
}

type AddressRepository interface {
	FindBySalonId(salonId string) ([]*Address, error)
	FindById(id string, salonId string) (*Address, error)
	Save(address *Address) (*Address, error)
	Delete(addressId string) error
}
