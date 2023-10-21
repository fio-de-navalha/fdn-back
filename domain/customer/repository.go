package customer

type CustomerRepository interface {
	FindMany() ([]*Customer, error)
	FindById(id string) (*Customer, error)
	FindByPhone(phone string) (*Customer, error)
	Save(customer *Customer) (*Customer, error)
}
