package customer

type CustomerRepository interface {
	FindMany() ([]*Customer, error)
	FindById(id string) (*Customer, error)
	FindByPhone(phone string) (*Customer, error)
	Create(customer *Customer) (*Customer, error)
}
