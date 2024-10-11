package auth_data

import (
	"aprendiendoGo/app/services/db/mongo"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Nombre de la colección de usuarios
const nameCollection = "users"

// CreateUser crea un nuevo usuario en la base de datos
func Create(user User) (User, error) {
	client := mongo.Open()
	collection := client.Database("examples").Collection(nameCollection)

	// Establecer el contexto con un timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Insertar el usuario en la colección
	_, err := collection.InsertOne(ctx, bson.M{
		"given_name":   user.GivenName,
		"family_name":  user.FamilyName,
		"email":        user.Email,
		"password":     user.Password, // Aquí se debe hashear la contraseña (luego se hará)
		"document":     user.Document,
		"documenttype": user.DocumentType,
		"photo":        user.Photo,
		"bio":          user.Bio,
		"birthday":     user.Birthday,
		"phone":        user.Phone,
		"created_at":   time.Now(),
		"updated_at":   time.Now(),
	})

	if err != nil {
		return User{}, err
	}

	// Retornar el usuario creado
	return user, nil
}
