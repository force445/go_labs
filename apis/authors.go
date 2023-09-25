package apis

import (
	"context"
	"fmt"
	databases "labs/database"
	"labs/models"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var authorCollection = databases.GetDB()
var validate = validator.New()

func CreateAuthor(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	var author models.Author
	defer cancel()

	if err := c.BodyParser(&author); err != nil {
		fmt.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http.StatusBadRequest, "message": "Error", "data": err.Error()})
	}

	if validationErr := validate.Struct(author); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": http.StatusBadRequest, "message": "Error", "data": validationErr.Error()})
	}

	newAuthor := models.Author{
		ID:   primitive.NewObjectID(),
		Name: author.Name,
		Age:  author.Age,
	}

	result, err := authorCollection.InsertOne(ctx, newAuthor)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": http.StatusInternalServerError, "message": "Error", "data": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"status": http.StatusCreated, "message": "Success", "data": result.InsertedID})
}

func GetAAuthor(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	id := c.Params("ID")
	var author models.Author
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)

	err := authorCollection.FindOne(ctx, bson.M{"id": objID}).Decode(&author)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": http.StatusInternalServerError, "message": "Error", "data": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"status": http.StatusOK, "message": "Success", "data": author})
}

// func UpdateAuthor(c *fiber.Ctx) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	id := c.Params("ID")
// 	var author models.Author
// 	defer cancel()

// }
