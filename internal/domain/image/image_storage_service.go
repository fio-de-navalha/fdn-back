package image

import "mime/multipart"

type ImageResponse struct {
	ID       string   `json:"id"`
	FileName string   `json:"filename"`
	Urls     []string `json:"urls"`
}

type ImageStorageService interface {
	GetImageById(imageId string) (*ImageResponse, error)
	UploadImage(fileBuffer *multipart.FileHeader) (*ImageResponse, error)
	UpdateImage(imageId string, fileBuffer *multipart.FileHeader) (*ImageResponse, error)
	DeleteImage(imageId string) error
}
