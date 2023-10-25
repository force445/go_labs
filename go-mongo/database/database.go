package databases

import (
	"context"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LoadEnv() string {
	godotenv.Load()

	return os.Getenv("MONGO_URL")
}

func ConnectMongoDB() (*mongo.Client, error) {
	url := LoadEnv()

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(url)

	client, err := mongo.Connect(context.Background(), clientOptions)

	return client, err
}

func GetDB() *mongo.Collection {
	client, _ := ConnectMongoDB()
	db := os.Getenv("MONGO_DB")
	collection := os.Getenv("MONGO_COLLECTION")

	return client.Database(db).Collection(collection)
}

func DisconnectFromMongo(client *mongo.Client) error {
	err := client.Disconnect(context.Background())
	return err
}
