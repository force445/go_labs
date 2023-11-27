package handler

import (
	"go-posgres/database"
	"go-posgres/model"

	"github.com/gofiber/fiber/v2"
)

func PostMember(c *fiber.Ctx) error {
	db := database.DB.Db
	member := new(model.Member)
	if err := c.BodyParser(member); err != nil {
		return err
	}
	db.Create(&member)
	return c.JSON(member)
}

func GetMember(c *fiber.Ctx) error {
	db := database.DB.Db
	member := new(model.Member)
	id := c.Params("id")
	db.First(&member, id)
	return c.JSON(member)
}

func GetAllMember(c *fiber.Ctx) error {
	db := database.DB.Db
	var members []model.Member
	db.Find(&members)
	return c.JSON(members)
}

func UpdateMember(c *fiber.Ctx) error {
	db := database.DB.Db
	member := new(model.Member)
	id := c.Params("id")
	db.First(&member, id)
	if err := c.BodyParser(member); err != nil {
		return err
	}
	db.Save(&member)
	return c.JSON(member)
}

func DeleteMember(c *fiber.Ctx) error {
	db := database.DB.Db
	member := new(model.Member)
	id := c.Params("id")
	db.First(&member, id)
	db.Delete(&member)
	return c.JSON(member)
}

func PostCreditCard(c *fiber.Ctx) error {
	db := database.DB.Db
	creditCard := new(model.CreditCard)
	if err := c.BodyParser(creditCard); err != nil {
		return err
	}
	db.Create(&creditCard)
	return c.JSON(creditCard)
}
