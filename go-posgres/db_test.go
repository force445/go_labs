package main

import (
	"fmt"
	"testing"

	"go-posgres/config"

	"github.com/joho/godotenv"
)

func TestConfig(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	expected := "postgres"
	actual := config.Config("DB_NAME")
	fmt.Println("actual: ", actual)
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}
