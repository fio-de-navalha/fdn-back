package product

type ProductRepository interface {
	FindManyByIds(ids []string) ([]*Product, error)
	FindById(id string) (*Product, error)
	FindBySalonId(salonId string) ([]*Product, error)
	Save(product *Product) (*Product, error)
}
