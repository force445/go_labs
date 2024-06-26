package handler

import (
	"go-posgres/database"
	"go-posgres/model"
	"go-posgres/validator"

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
	db.Preload("CreditCards").Preload("Orders").Find(&members)
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

	isValid, validationErrors := validator.ValidateCreditCard(*creditCard)
	if !isValid {
		return c.Status(400).JSON(fiber.Map{"errors": validationErrors})
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
	db.Preload("Books")
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

func PostBook(c *fiber.Ctx) error {
	db := database.DB.Db
	book := new(model.Book)
	if err := c.BodyParser(book); err != nil {
		return err
	}
	db.Create(&book)
	return c.JSON(book)
}

func GetBook(c *fiber.Ctx) error {
	db := database.DB.Db
	book := new(model.Book)
	id := c.Params("id")
	db.First(&book, id)
	return c.JSON(book)
}

func GetAllBook(c *fiber.Ctx) error {
	db := database.DB.Db
	var books []model.Book
	db.Find(&books)
	return c.JSON(books)
}

func UpdateBook(c *fiber.Ctx) error {
	db := database.DB.Db
	book := new(model.Book)
	id := c.Params("id")
	db.First(&book, id)
	if err := c.BodyParser(book); err != nil {
		return err
	}
	db.Save(&book)
	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	db := database.DB.Db
	book := new(model.Book)
	id := c.Params("id")
	db.First(&book, id)
	db.Delete(&book)
	return c.JSON(book)
}

func PostAuthor(c *fiber.Ctx) error {
	db := database.DB.Db
	author := new(model.Author)
	if err := c.BodyParser(author); err != nil {
		return err
	}
	db.Create(&author)
	return c.JSON(author)
}

func GetAuthor(c *fiber.Ctx) error {
	db := database.DB.Db
	author := new(model.Author)
	id := c.Params("id")
	db.First(&author, id)
	db.Preload("Books")
	return c.JSON(author)
}

func UpdateAuthor(c *fiber.Ctx) error {
	db := database.DB.Db
	author := new(model.Author)
	id := c.Params("id")
	db.First(&author, id)
	if err := c.BodyParser(author); err != nil {
		return err
	}
	db.Save(&author)
	return c.JSON(author)
}

func DeleteAuthor(c *fiber.Ctx) error {
	db := database.DB.Db
	author := new(model.Author)
	id := c.Params("id")
	db.First(&author, id)
	db.Delete(&author)
	return c.JSON(author)
}

func PostPublisher(c *fiber.Ctx) error {
	db := database.DB.Db
	publisher := new(model.Publisher)
	if err := c.BodyParser(publisher); err != nil {
		return err
	}
	db.Create(&publisher)
	return c.JSON(publisher)
}

func GetPublisher(c *fiber.Ctx) error {
	db := database.DB.Db
	publisher := new(model.Publisher)
	id := c.Params("id")
	db.First(&publisher, id)
	db.Preload("Books")
	return c.JSON(publisher)
}

func UpdatePublisher(c *fiber.Ctx) error {
	db := database.DB.Db
	publisher := new(model.Publisher)
	id := c.Params("id")
	db.First(&publisher, id)
	if err := c.BodyParser(publisher); err != nil {
		return err
	}
	db.Save(&publisher)
	return c.JSON(publisher)
}

func DeletePublisher(c *fiber.Ctx) error {
	db := database.DB.Db
	publisher := new(model.Publisher)
	id := c.Params("id")
	db.First(&publisher, id)
	db.Delete(&publisher)
	return c.JSON(publisher)
}
