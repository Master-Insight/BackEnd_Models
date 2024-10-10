package auth_data

import (
	"time"
)

// User representa la estructura de un usuario
type User struct {
	GivenName    string    `json:"given_name" validate:"required,max=50" field:"Nombre" required-message:"El nombre es obligatorio"`
	FamilyName   string    `json:"family_name" validate:"required,max=50" field:"Apellido"`
	Email        string    `json:"email" validate:"required,email" field:"Correo electrónico"`
	Password     string    `json:"password" validate:"required,min=6" field:"Contraseña"`
	Document     string    `json:"document" validate:"max=15" field:"Documento"`
	DocumentType string    `json:"documenttype" validate:"omitempty,oneof=DNI LC PASSPORT" field:"Tipo de documento"`
	Photo        string    `json:"photo" validate:"omitempty,url" field:"Foto"`
	Bio          string    `json:"bio" validate:"omitempty" field:"Biografía"`
	Birthday     time.Time `json:"birthday" validate:"omitempty" field:"Fecha de nacimiento"`
	Phone        string    `json:"phone" validate:"omitempty,max=20" field:"Teléfono"`
}
