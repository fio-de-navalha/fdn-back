package barber

type Contact struct {
	ID       uint   `json:"id"`
	BarberId string `json:"barber_id"`
	Contact  string `json:"contact"`
}

func NewContact(barberId string, contact string) *Contact {
	return &Contact{
		BarberId: barberId,
		Contact:  contact,
	}
}
