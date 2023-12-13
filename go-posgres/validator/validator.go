package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"

	"go-posgres/model"
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

func ValidateCreditCard(creditCardData model.CreditCard) (bool, []string) {
	validate := validator.New()
	err := validate.Struct(creditCardData)

	if err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
		}
		return false, errors
	}

	if !luhnAlgorithm(creditCardData.CardNo) {
		return false, []string{"Invalid credit card number"}
	}

	return true, nil
}

func luhnAlgorithm(cardNumber string) bool {
	var sum int
	var alternate bool = false

	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit := int(cardNumber[i] - '0')

		if alternate {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		alternate = !alternate
	}

	return sum%10 == 0
}
