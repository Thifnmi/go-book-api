package domain

type MetadataResponse struct {
	Total       int `json:"total"`
	CurrentPage int `json:"current_page"`
	Pages       int `json:"pages"`
}

type Response struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	ErrorCode int16       `json:"error_code"`
	Data      interface{} `json:"data"`
}
