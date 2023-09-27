package cloudflare

import "time"

type imageResult struct {
	ID                string    `json:"id"`
	FileName          string    `json:"filename"`
	Uploaded          time.Time `json:"uploaded"`
	RequireSignedURLs bool      `json:"requireSignedURLs"`
	Variants          []string  `json:"variants"`
}

type singleImageResponse struct {
	Result   imageResult   `json:"result"`
	Success  bool          `json:"success"`
	Errors   []interface{} `json:"errors"`
	Messages []interface{} `json:"messages"`
}

type imagesResult struct {
	Images []imageResult `json:"images"`
}

type multipleImagesResponse struct {
	Result   imagesResult  `json:"result"`
	Success  bool          `json:"success"`
	Errors   []interface{} `json:"errors"`
	Messages []interface{} `json:"messages"`
}
