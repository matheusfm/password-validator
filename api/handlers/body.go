package handlers

type ValidationRequest struct {
	Password string `json:"password"`
}

type ValidationResponse struct {
	Valid   bool   `json:"valid"`
	Message string `json:"message,omitempty"`
}
