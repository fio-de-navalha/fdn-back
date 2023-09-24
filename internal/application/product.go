package application

import (
	"errors"
	"mime/multipart"

	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/image"
	"github.com/fio-de-navalha/fdn-back/internal/domain/product"
)

type ProductService struct {
	productRepository   product.ProductRepository
	barberService       BarberService
	imageStorageService image.ImageStorageService
}

func NewProductService(
	productRepository product.ProductRepository,
	barberService BarberService,
	imageStorageService image.ImageStorageService,
) *ProductService {
	return &ProductService{
		productRepository:   productRepository,
		barberService:       barberService,
		imageStorageService: imageStorageService,
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

func (s *ProductService) CreateProduct(input product.CreateProductRequest, file *multipart.FileHeader) error {
	barberExists, err := s.barberService.GetBarberById(input.BarberId)
	if err != nil {
		return err
	}
	if barberExists == nil {
		return errors.New("barber not found")
	}

	if file != nil {
		file.Filename = constants.FilePrefix + file.Filename
		res, err := s.imageStorageService.UploadImage(file)
		if err != nil {
			return err
		}

		input.ImageId = res.ID
		input.ImageUrl = res.Urls[0]
	}

	ser := product.NewProduct(input)
	_, err = s.productRepository.Save(ser)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductService) UpdateProduct(productId string, input product.UpdateProductRequest) error {
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
