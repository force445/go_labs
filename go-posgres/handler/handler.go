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
	db.Preload("CreditCards").First(&member, id)

	return c.JSON(member)
}

func GetAllMember(c *fiber.Ctx) error {
	db := database.DB.Db
	var members []model.Member
	db.Preload("CreditCards").Find(&members)
	db.Preload("Orders").Find(&members)
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

func GetCreditCard(c *fiber.Ctx) error {
	db := database.DB.Db
	creditcard := new(model.CreditCard)
	id := c.Params("id")
	db.First(&creditcard, id)
	return c.JSON(creditcard)
}

func UpdateCreditCard(c *fiber.Ctx) error {
	db := database.DB.Db
	creditcard := new(model.CreditCard)
	id := c.Params("id")
	db.First(&creditcard, id)
	if err := c.BodyParser(creditcard); err != nil {
		return err
	}
	db.Save(&creditcard)
	return c.JSON(creditcard)
}

func DeleteCreditCard(c *fiber.Ctx) error {
	db := database.DB.Db
	creditcard := new(model.CreditCard)
	id := c.Params("id")
	db.First(&creditcard, id)
	db.Delete(&creditcard)
	return c.JSON(creditcard)
}

func PostOrder(c *fiber.Ctx) error {
	db := database.DB.Db
	order := new(model.Order)
	if err := c.BodyParser(order); err != nil {
		return err
	}
	db.Create(&order)
	return c.JSON(order)
}

func GetOrder(c *fiber.Ctx) error {
	db := database.DB.Db
	order := new(model.Order)
	id := c.Params("id")
	db.First(&order, id)
	return c.JSON(order)
}

func GetAllOrder(c *fiber.Ctx) error {
	db := database.DB.Db
	var orders []model.Order
	db.Find(&orders)
	return c.JSON(orders)
}

func UpdateOrder(c *fiber.Ctx) error {
	db := database.DB.Db
	order := new(model.Order)
	id := c.Params("id")
	db.First(&order, id)
	if err := c.BodyParser(order); err != nil {
		return err
	}
	db.Save(&order)
	return c.JSON(order)
}

func DeleteOrder(c *fiber.Ctx) error {
	db := database.DB.Db
	order := new(model.Order)
	id := c.Params("id")
	db.First(&order, id)
	db.Delete(&order)
	return c.JSON(order)
}
