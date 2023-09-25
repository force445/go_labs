package main

import (
	"context"
	databases "labs/database"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestEnv(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Error(err)
	}
}

func TestMongoConnection(t *testing.T) {
	_, err := databases.ConnectMongoDB()

	if err != nil {
		t.Error(err)
	}
}

func TestConnectionTimeout(t *testing.T) {
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, _ := databases.ConnectMongoDB()

	err := client.Ping(context.Background(), nil)

	if err != nil {
		t.Error(err)
	}

}

func TestDB(t *testing.T) {
	db := databases.GetDB()

	if db == nil {
		t.Error("DB is nil")
	}
}

// func TestCreateAuthor(t *testing.T) {
// 	app := fiber.New()
// 	app.Post("/Author", apis.CreateAuthor)

// 	req, _ := app.Test(httptest.NewRequest("POST", "/Author", strings.NewReader(`{"name":"test","age":20}`)))

// 	if req.StatusCode != 201 {
// 		t.Error("Status:", req.StatusCode)
// 	}

// }

func TestDisconnectMongo(t *testing.T) {
	client, _ := databases.ConnectMongoDB()
	err := databases.DisconnectFromMongo(client)
	if err != nil {
		t.Error(err)
	}
}
