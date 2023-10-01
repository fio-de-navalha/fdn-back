package application

import (
	"errors"
	"log"
	"mime/multipart"

	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/image"
	"github.com/fio-de-navalha/fdn-back/internal/domain/product"
)

type ProductService struct {
	productRepository   product.ProductRepository
	salonService        SalonService
	professionalService ProfessionalService
	imageStorageService image.ImageStorageService
}

func NewProductService(
	productRepository product.ProductRepository,
	salonService SalonService,
	professionalService ProfessionalService,
	imageStorageService image.ImageStorageService,
) *ProductService {
	return &ProductService{
		productRepository:   productRepository,
		salonService:        salonService,
		professionalService: professionalService,
		imageStorageService: imageStorageService,
	}
}

func (s *ProductService) GetProductsBySalonId(salonId string) ([]*product.Product, error) {
	log.Println("[ProductService.GetProductsBySalonId] - Validating salon:", salonId)
	sal, err := s.salonService.GetSalonById(salonId)
	if err != nil {
		return nil, err
	}
	if sal == nil {
		return nil, errors.New("salon not found")
	}

	log.Println("[ProductService.GetProductsBySalonId] - Getting products from salon:", salonId)
	res, err := s.productRepository.FindBySalonId(salonId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ProductService) CreateProduct(input product.CreateProductRequest, file *multipart.FileHeader) error {
	log.Println("[ProductService.CreateProduct] - Validating salon:", input.SalonId)
	sal, err := s.salonService.validateSalon(input.SalonId)
	if err != nil {
		return err
	}

	log.Println("[ProductService.CreateProduct] - Validating professional:", input.ProfessionalId)
	prof, err := s.professionalService.validateProfessionalById(input.ProfessionalId)
	if err != nil {
		return err
	}

	if err := s.salonService.validateRequesterPermission(prof.ID, sal.SalonMembers); err != nil {
		return err
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

	log.Println("[ProductService.CreateProduct] - Creating product")
	newProduct := product.NewProduct(input)
	_, err = s.productRepository.Save(newProduct)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductService) UpdateProduct(productId string, input product.UpdateProductRequest, file *multipart.FileHeader) error {
	log.Println("[ProductService.UpdateProduct] - Validating salon:", input.SalonId)
	sal, err := s.salonService.validateSalon(input.SalonId)
	if err != nil {
		return err
	}

	log.Println("[ProductService.UpdateProduct] - Validating professional:", input.ProfessionalId)
	prof, err := s.professionalService.validateProfessionalById(input.ProfessionalId)
	if err != nil {
		return err
	}

	if err := s.salonService.validateRequesterPermission(prof.ID, sal.SalonMembers); err != nil {
		return err
	}

	log.Println("[ProductService.UpdateProduct] - Validating product:", productId)
	pro, err := s.validateProduct(productId)
	if err != nil {
		return err
	}

	if file != nil {
		log.Println("[ProductService.UpdateProduct] - Updating image")
		res, err := s.imageStorageService.UpdateImage(pro.ImageId, file)
		if err != nil {
			return err
		}

		pro.ImageId = res.ID
		pro.ImageUrl = res.Urls[0]
	}

	if input.Name != nil {
		pro.Name = *input.Name
	}
	if input.Price != nil {
		pro.Price = *input.Price
	}
	if input.Available != nil {
		pro.Available = *input.Available
	}

	log.Println("[ProductService.UpdateProduct] - Updating product")
	if _, err = s.productRepository.Save(pro); err != nil {
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

func (s *ProductService) validateProductsAvailability(products []*product.Product) []string {
	var productsIdsToSave []string
	for _, v := range products {
		if v.Available {
			productsIdsToSave = append(productsIdsToSave, v.ID)
		}
	}
	return productsIdsToSave
}

func (s *ProductService) validateProduct(productId string) (*product.Product, error) {
	pro, err := s.productRepository.FindById(productId)
	if err != nil {
		return nil, err
	}
	if pro == nil {
		return nil, errors.New("product not found")
	}
	return pro, nil
}
