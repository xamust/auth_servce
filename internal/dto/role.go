package dto

type Role struct {
	Description string       `json:"description"`
	Title       string       `json:"title"`
	Permissions []Permission `json:"permissions"`
}

type Permission struct {
	Description string `json:"description"`
	Title       string `json:"title"`
}
