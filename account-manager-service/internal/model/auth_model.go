package model

type SignUpRequest struct {
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Password    string `json:"password" validate:"required,min=5"`
	PhoneNumber string `json:"phone_number" validate:"required,e164"`
	Email       string `json:"email" validate:"required,email"`
}
