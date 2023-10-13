package salon

type SalonMemberRepository interface {
	FindBySalonId(salonId string) ([]*SalonMember, error)
	FindById(id string, salonId string) (*SalonMember, error)
	FindByProfessionalId(professionalIdid string) (*SalonMember, error)
	Save(salonMember *SalonMember) (*SalonMember, error)
	Delete(salonMemberId string) error
}
