package barber

type AddressRepository interface {
	FindMany() ([]*Address, error)
	FindById(id string) (*Address, error)
	Save(address *Address) (*Address, error)
}
