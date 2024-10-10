package validatorpartner

// Custom error messages for validation tags
var customErrorMessages = map[string]string{
	"GivenName.required":  "El nombre es obligatorio",
	"GivenName.max":       "El nombre no debe exceder los 50 caracteres",
	"FamilyName.required": "El apellido es obligatorio",
	"FamilyName.max":      "El apellido no debe exceder los 50 caracteres",
	"Email.required":      "El correo es obligatorio",
	"Email.email":         "El formato del correo no es válido",
	"Password.required":   "La contraseña es obligatoria",
	"Password.min":        "La contraseña debe tener al menos 6 caracteres",
	"Document.max":        "El documento no puede exceder los 15 caracteres",
	"DocumentType.oneof":  "El tipo de documento debe ser DNI, LC o PASSPORT",
	"Photo.url":           "La URL de la foto no es válida",
	"Phone.max":           "El teléfono no puede exceder los 20 caracteres",
}
