package product

type ProductRepository interface {
	FindManyByIds(ids []uint) ([]*Product, error)
	FindById(id uint) (*Product, error)
	FindByBarberId(barberId string) ([]*Product, error)
	Save(product *Product) (*Product, error)
}
