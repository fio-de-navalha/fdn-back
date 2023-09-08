package application

import (
	"errors"

	"github.com/fio-de-navalha/fdn-back/internal/domain/product"
)

type ProductService struct {
	productRepository product.ProductRepository
	barberService     BarberService
}

func NewProductService(productRepository product.ProductRepository, barberService BarberService) *ProductService {
	return &ProductService{
		productRepository: productRepository,
		barberService:     barberService,
	}
}

func (s *ProductService) GetProductsByBarberId(barberId string) ([]*product.Product, error) {
	barberExists, err := s.barberService.GetBarberById(barberId)
	if err != nil {
		return nil, err
	}
	if barberExists == nil {
		return nil, errors.New("barber not found")
	}

	res, err := s.productRepository.FindByBarberId(barberId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ProductService) CreateProduct(input product.CreateProductInput) error {
	barberExists, err := s.barberService.GetBarberById(input.BarberId)
	if err != nil {
		return err
	}
	if barberExists == nil {
		return errors.New("barber not found")
	}

	ser := product.NewProduct(input)
	_, err = s.productRepository.Save(ser)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductService) UpdateProduct(productId string, input product.UpdateProductInput) error {
	ser, err := s.productRepository.FindById(productId)
	if err != nil {
		return err
	}
	if ser == nil {
		return errors.New("product not found")
	}

	if input.Name != nil {
		ser.Name = *input.Name
	}
	if input.Price != nil {
		ser.Price = *input.Price
	}
	if input.Available != nil {
		ser.Available = *input.Available
	}

	_, err = s.productRepository.Save(ser)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductService) getManyProducts(productIds []string) ([]*product.Product, error) {
	res, err := s.productRepository.FindManyByIds(productIds)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ProductService) ValidateProductsAvailability(products []*product.Product) []string {
	var productsIdsToSave []string
	for _, v := range products {
		if v.Available {
			productsIdsToSave = append(productsIdsToSave, v.ID)
		}
	}

	return productsIdsToSave
}
