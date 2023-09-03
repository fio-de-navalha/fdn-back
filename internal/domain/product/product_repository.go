package product

type ProductRepository interface {
	FindById(id uint) (*Product, error)
	FindByBarberId(barberId string) ([]*Product, error)
	Save(product *Product) (*Product, error)
}
