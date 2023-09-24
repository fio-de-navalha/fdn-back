package cloudflare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/fio-de-navalha/fdn-back/internal/constants"
	"github.com/fio-de-navalha/fdn-back/internal/domain/image"
)

type CloudFlareService struct {
	baseUrl   string
	accountId string
	readToken string
	editToken string
}

func NewCloudFlareService(
	baseUrl string,
	accountId string,
	readToken string,
	editToken string,
) *CloudFlareService {
	return &CloudFlareService{
		baseUrl:   baseUrl,
		accountId: accountId,
		readToken: readToken,
		editToken: editToken,
	}
}

func (s *CloudFlareService) GetImageById(imageId string) (*image.ImageResponse, error) {
	log.Println("[cloudflare.GetImageById] - Getting image:", imageId)
	url := s.baseUrl + "/" + imageId
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	token := "Bearer " + s.readToken
	req.Header.Add("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var response singleImageResponse
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil, err
	}

	return &image.ImageResponse{
		ID:       response.Result.ID,
		FileName: response.Result.FileName,
		Urls:     response.Result.Variants,
	}, nil
}

func (s *CloudFlareService) UploadImage(file *multipart.FileHeader) (*image.ImageResponse, error) {
	log.Println("[cloudflare.UploadImage] - Upload image:", file.Filename)
	fileContent, err := file.Open()
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer fileContent.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	filePart, err := writer.CreateFormFile("file", file.Filename)
	if err != nil {
		fmt.Println("Error creating form file:", err)
		return nil, err
	}

	if _, err = io.Copy(filePart, fileContent); err != nil {
		fmt.Println("Error copying file to form field:", err)
		return nil, err
	}
	writer.Close()

	req, err := http.NewRequest("POST", s.baseUrl, body)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	token := "Bearer " + s.editToken
	req.Header.Add("Authorization", token)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var response singleImageResponse
	if err := json.Unmarshal(resBody, &response); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil, err
	}

	return &image.ImageResponse{
		ID:       response.Result.ID,
		FileName: response.Result.FileName,
		Urls:     response.Result.Variants,
	}, nil
}

func (s *CloudFlareService) UpdateImage(imageId string, file *multipart.FileHeader) (*image.ImageResponse, error) {
	if file == nil {
		return nil, nil
	}

	deleteErrorCh := make(chan error)
	uploadErrorCh := make(chan error)
	uploadResultCh := make(chan *image.ImageResponse)

	go func() {
		err := s.DeleteImage(imageId)
		deleteErrorCh <- err
	}()
	go func() {
		file.Filename = constants.FilePrefix + file.Filename
		res, err := s.UploadImage(file)
		uploadErrorCh <- err
		uploadResultCh <- res
	}()

	deleteErr := <-deleteErrorCh
	uploadErr := <-uploadErrorCh
	if deleteErr != nil {
		return nil, deleteErr
	}
	if uploadErr != nil {
		return nil, uploadErr
	}

	uploaded := <-uploadResultCh

	return uploaded, nil
}

func (s *CloudFlareService) DeleteImage(imageId string) error {
	log.Println("[cloudflare.GetImageById] - Getting image:", imageId)
	url := s.baseUrl + "/" + imageId
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	token := "Bearer " + s.editToken
	req.Header.Add("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	return nil
}
