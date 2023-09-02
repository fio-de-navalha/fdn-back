package product

type ProductRepository interface {
	FindById(id string) (*Product, error)
	FindByBarberId(barberId string) ([]*Product, error)
	Save(product *Product) (*Product, error)
}
