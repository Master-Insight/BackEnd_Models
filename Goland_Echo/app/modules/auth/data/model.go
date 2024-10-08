package auth_data

import "time"

// User representa la estructura de un usuario
type User struct {
	GivenName    string    `json:"given_name" validate:"required,max=50"`
	FamilyName   string    `json:"family_name" validate:"required,max=50"`
	Email        string    `json:"email" validate:"required,email"`
	Password     string    `json:"password" validate:"required,min=6"`
	Document     string    `json:"document" validate:"max=15"`
	DocumentType string    `json:"documenttype" validate:"omitempty,oneof=DNI LC PASSPORT"`
	Photo        string    `json:"photo" validate:"omitempty,url"`
	Bio          string    `json:"bio" validate:"omitempty"`
	Birthday     time.Time `json:"birthday" validate:"omitempty"`
	Phone        string    `json:"phone" validate:"omitempty,max=20"`
}
