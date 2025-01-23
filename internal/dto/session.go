package dto

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	ID  string `json:"id"`
	TTL int64  `json:"ttl"`
}
