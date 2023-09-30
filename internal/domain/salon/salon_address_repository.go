package salon

type AddressRepository interface {
	FindBySalonId(salonId string) ([]*Address, error)
	FindById(id string, salonId string) (*Address, error)
	Save(address *Address) (*Address, error)
	Delete(addressId string) error
}
