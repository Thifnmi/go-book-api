package domain

type HealthResponse struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type InfoResponse struct {
	Auth     string `json:"Auth"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Telegram string `json:"telegram"`
}
