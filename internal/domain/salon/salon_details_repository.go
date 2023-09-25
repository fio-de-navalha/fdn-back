package salon

type SalonMemberRepository interface {
	FindBySalonId(salonId string) ([]*SalonMember, error)
	FindById(id string, salonId string) (*SalonMember, error)
	Save(salonMember *SalonMember) (*SalonMember, error)
	Delete(salonMemberId string) error
}

type AddressRepository interface {
	FindBySalonId(salonId string) ([]*Address, error)
	FindById(id string, salonId string) (*Address, error)
	Save(address *Address) (*Address, error)
	Delete(addressId string) error
}

type ContactRepository interface {
	FindBySalonId(salonId string) ([]*Contact, error)
	FindById(id string, salonId string) (*Contact, error)
	Save(contact *Contact) (*Contact, error)
	Delete(contactId string) error
}
