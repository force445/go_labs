package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"

	"go-posgres/handler"
)

func ValidateStruct(data interface{}) (bool, []string) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	validate := validator.New()
	err = validate.Struct(data)

	if err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
		}
		return false, errors
	}

	return true, nil
}

func ValidateCreditCard() {
	creditcard := handler.GetCreditCard()
}
