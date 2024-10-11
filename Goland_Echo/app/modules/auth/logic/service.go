package auth_logic

import auth_data "aprendiendoGo/app/modules/auth/data"

// RegisterUser procesa los datos del usuario y los devuelve (placeholder)
func RegisterUser(user auth_data.User) (auth_data.User, error) {

	// Llamar al DAO para crear el usuario en MongoDB
	createdUser, err := auth_data.Create(user)
	if err != nil {
		return auth_data.User{}, err
	}

	// Devolver el usuario creado
	return createdUser, nil
}
