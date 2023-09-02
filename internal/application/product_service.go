package application

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/fio-de-navalha/fdn-back/internal/domain/product"
)

type ProductService struct {
	productRepository product.ProductRepository
	barberService     BarberService
}

func NewProductService(productRepository product.ProductRepository, barberService BarberService) *ProductService {
	return &ProductService{
		productRepository: productRepository,
		barberService: barberService,
	}
}

func (s *ProductService) GetProductsByBarberId(barberId string) ([]*product.Product, error){
	barberExists, err := s.barberService.GetBarberById(barberId)
	if err != nil {
		return nil, err
	}
	if barberExists == nil {
		return nil, errors.New("barber not found")
	}

	res, err := s.productRepository.FindByBarberId(barberId)
	if err != nil {
		// TODO: add better error handling
		fmt.Println(err)
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
		// TODO: add better error handling
		fmt.Println(err)
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

	updateField := func(dest, source interface{}) {
		if source != nil {
			reflect.ValueOf(dest).Elem().Set(reflect.ValueOf(source).Elem())
		}
	}
	updateField(&ser.Name, input.Name)
	updateField(&ser.Price, input.Price)
	updateField(&ser.Available, input.Available)

	_, err = s.productRepository.Save(ser)
	if err != nil {
		// TODO: add better error handling
		fmt.Println(err)
	}
	return nil
}