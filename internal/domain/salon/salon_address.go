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
}

func NewAddress(salonId string, address string) *Address {
	return &Address{
		ID:      uuid.NewString(),
		SalonId: salonId,
		Address: address,
	}
}
