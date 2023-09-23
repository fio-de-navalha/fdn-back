package product

type ProductRepository interface {
	FindManyByIds(ids []string) ([]*Product, error)
	FindById(id string) (*Product, error)
	FindByBarberId(barberId string) ([]*Product, error)
	Save(product *Product) (*Product, error)
}
