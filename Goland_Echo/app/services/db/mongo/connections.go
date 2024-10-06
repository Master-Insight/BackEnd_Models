package mongo

import (
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoSingleton struct {
	client *mongo.Client
}

var instance *MongoSingleton
var once sync.Once

// Conectar a MongoDB (singleton)
func (m *MongoSingleton) connect(uri string) *mongo.Client {

	clientOptions := options.Client().ApplyURI(uri)

	// Crear el contexto con un timeout para la conexión
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Error al conectar a MongoDB: %v", err)
	}

	// Verificar la conexión
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("Error al hacer ping a MongoDB: %v", err)
	}

	log.Println("Conectado a MongoDB")
	return client
}

// Obtener una instancia única de la conexión a MongoDB
func GetMongoInstance(uri string) *mongo.Client {
	once.Do(func() {
		instance = &MongoSingleton{}
		instance.client = instance.connect(uri)
	})

	return instance.client
}

// Cerrar la conexión cuando sea necesario
func DisconnectMongo() {
	if instance != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := instance.client.Disconnect(ctx)
		if err != nil {
			log.Fatalf("Error al desconectar de MongoDB: %v", err)
		}

		log.Println("Conexión a MongoDB cerrada")
	}
}
