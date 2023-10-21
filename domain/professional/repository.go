package professional

type ProfessionalRepository interface {
	FindMany() ([]*Professional, error)
	FindById(id string) (*Professional, error)
	FindByEmail(email string) (*Professional, error)
	Save(professional *Professional) (*Professional, error)
}
