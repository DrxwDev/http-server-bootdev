package chirpy

type ChirpRequestDTO struct {
	Body   string `json:"body"`
	UserID string `json:"user_id"`
}

type ChirpResponseDTO struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Body      string `json:"body"`
	UserID    string `json:"user_id"`
}
