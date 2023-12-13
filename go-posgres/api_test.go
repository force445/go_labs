package main

import (
	"bytes"
	"fmt"
	"net/http/httptest"
	"testing"

	"go-posgres/config"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	expected := "demo"
	actual := config.Config("DB_NAME")
	fmt.Println("actual: ", actual)
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestGetRoute(t *testing.T) {
	app := fiber.New()
	app.Get("/api/v1/member/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	req := httptest.NewRequest("GET", "/api/v1/member/", nil)
	resp, err := app.Test(req, 1000)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 200, resp.StatusCode)
}

func TestPostRoute(t *testing.T) {
	app := fiber.New()

	requestPayload := []byte(`{"Fullname": "John Doe","Birthday": "testtest","Address": "testtesttest","Province": "ooooooo","City": "testtesttest","Zipcode": "testtesttest","Email": "testtesttest","Phone": "testtesttest"}`)

	req := httptest.NewRequest("POST", "/api/v1/member/", bytes.NewBuffer(requestPayload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, 1000)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, resp.StatusCode)
}
