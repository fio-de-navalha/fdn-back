package barber

type Address struct {
	BarberId string `json:"barber_id"`
	Address  string `json:"address"`
}

func NewAddress(barberId string, address string) *Address {
	return &Address{
		BarberId: barberId,
		Address:  address,
	}
}
