package product

type Product struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	BarberId  string `json:"barberId"`
	Name      string `json:"name"`
	Price     int32  `json:"price"`
	Available bool   `json:"available"`
}

func NewProduct(input CreateProductInput) *Product {
	return &Product{
		BarberId:  input.BarberId,
		Name:      input.Name,
		Price:     input.Price,
		Available: true,
	}
}
