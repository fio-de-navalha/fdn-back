package app

import (
	"log/slog"
	"mime/multipart"

	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/image"
	"github.com/fio-de-navalha/fdn-back/internal/domain/salon"
	"github.com/fio-de-navalha/fdn-back/pkg/errors"
)

type ProductService struct {
	productRepository   salon.ProductRepository
	salonService        SalonService
	professionalService ProfessionalService
	imageStorageService image.ImageStorageService
}

func NewProductService(
	productRepository salon.ProductRepository,
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

func (s *ProductService) GetProductsBySalonId(salonId string) ([]*salon.Product, error) {
	slog.Info("[ProductService.GetProductsBySalonId] - Validating salon: " + salonId)
	sal, err := s.salonService.GetSalonById(salonId)
	if err != nil {
		return nil, err
	}
	if sal == nil {
		return nil, &errors.AppError{
			Code:    constants.SALON_NOT_FOUND_ERROR_CODE,
			Message: constants.SALON_NOT_FOUND_ERROR_MESSAGE,
		}
	}

	slog.Info("[ProductService.GetProductsBySalonId] - Getting products from salon: " + salonId)
	res, err := s.productRepository.FindBySalonId(salonId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ProductService) CreateProduct(input salon.CreateProductRequest, file *multipart.FileHeader) error {
	slog.Info("[ProductService.CreateProduct] - Validating salon: " + input.SalonId)
	sal, err := s.salonService.validateSalon(input.SalonId)
	if err != nil {
		return err
	}

	slog.Info("[ProductService.CreateProduct] - Validating professional: " + input.ProfessionalId)
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

	slog.Info("[ProductService.CreateProduct] - Creating product")
	newProduct := salon.NewProduct(input)
	_, err = s.productRepository.Save(newProduct)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductService) UpdateProduct(productId string, input salon.UpdateProductRequest, file *multipart.FileHeader) error {
	slog.Info("[ProductService.UpdateProduct] - Validating salon: " + input.SalonId)
	sal, err := s.salonService.validateSalon(input.SalonId)
	if err != nil {
		return err
	}

	slog.Info("[ProductService.UpdateProduct] - Validating professional: " + input.ProfessionalId)
	prof, err := s.professionalService.validateProfessionalById(input.ProfessionalId)
	if err != nil {
		return err
	}

	if err := s.salonService.validateRequesterPermission(prof.ID, sal.SalonMembers); err != nil {
		return err
	}

	slog.Info("[ProductService.UpdateProduct] - Validating product: " + productId)
	pro, err := s.validateProduct(productId)
	if err != nil {
		return err
	}

	if file != nil {
		slog.Info("[ProductService.UpdateProduct] - Updating image")
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

	slog.Info("[ProductService.UpdateProduct] - Updating product")
	if _, err = s.productRepository.Save(pro); err != nil {
		return err
	}
	return nil
}

func (s *ProductService) getManyProducts(productIds []string) ([]*salon.Product, error) {
	res, err := s.productRepository.FindManyByIds(productIds)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ProductService) validateProductsAvailability(products []*salon.Product) []string {
	var productsIdsToSave []string
	for _, v := range products {
		if v.Available {
			productsIdsToSave = append(productsIdsToSave, v.ID)
		}
	}
	return productsIdsToSave
}

func (s *ProductService) validateProduct(productId string) (*salon.Product, error) {
	pro, err := s.productRepository.FindById(productId)
	if err != nil {
		return nil, err
	}
	if pro == nil {
		return nil, &errors.AppError{
			Code:    constants.PRODUCT_NOT_FOUND_ERROR_CODE,
			Message: constants.PRODUCT_NOT_FOUND_ERROR_MESSAGE,
		}
	}
	return pro, nil
}
