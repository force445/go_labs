package handler

import (
	"fmt"
	"go-posgres/database"
	"go-posgres/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func PostUser(c *fiber.Ctx) error {
	db := database.DB.Db

	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}

	db.Create(&user)

	return c.Status(fiber.StatusCreated).JSON(user)

}

func GetAllUser(c *fiber.Ctx) error {
	db := database.DB.Db

	var users []model.User

	err := db.Model(&model.User{}).Preload("Location").Find(&users).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot get users",
		})
	}
	return c.Status(fiber.StatusOK).JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	db := database.DB.Db

	id := c.Params("id")

	var user model.User

	db.Find(&user, "id = ?", id)

	if user.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)

}

func PostLocation(c *fiber.Ctx) error {
	db := database.DB.Db

	location := new(model.Location)

	if err := c.BodyParser(location); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}

	db.Create(&location)

	return c.Status(fiber.StatusCreated).JSON(location)

}
