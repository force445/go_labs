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

func GetAllLocation(c *fiber.Ctx) error {
	db := database.DB.Db

	var locations []model.Location

	err := db.Model(&model.Location{}).Find(&locations).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot get locations",
		})
	}
	return c.Status(fiber.StatusOK).JSON(locations)
}

func GetLocation(c *fiber.Ctx) error {
	db := database.DB.Db

	id := c.Params("id")

	var location model.Location

	db.Find(&location, "id = ?", id)

	if location.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Location not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(location)
}

func PostStation(c *fiber.Ctx) error {
	db := database.DB.Db

	station := new(model.Station)

	if err := c.BodyParser(station); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}

	db.Create(&station)

	return c.Status(fiber.StatusCreated).JSON(station)

}

func GetAllStation(c *fiber.Ctx) error {
	db := database.DB.Db

	var stations []model.Station

	err := db.Model(&model.Station{}).Find(&stations).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot get stations",
		})
	}
	return c.Status(fiber.StatusOK).JSON(stations)
}

func GetStation(c *fiber.Ctx) error {
	db := database.DB.Db

	id := c.Params("id")

	var station model.Station

	db.Find(&station, "id = ?", id)

	if station.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Station not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(station)
}

func PostDrone(c *fiber.Ctx) error {
	db := database.DB.Db

	drone := new(model.Drone)

	if err := c.BodyParser(drone); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}

	db.Create(&drone)

	return c.Status(fiber.StatusCreated).JSON(drone)

}

func GetAllDrone(c *fiber.Ctx) error {
	db := database.DB.Db

	var drones []model.Drone

	err := db.Model(&model.Drone{}).Find(&drones).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot get drones",
		})
	}
	return c.Status(fiber.StatusOK).JSON(drones)
}

func GetDrone(c *fiber.Ctx) error {
	db := database.DB.Db

	id := c.Params("id")

	var drone model.Drone

	db.Find(&drone, "id = ?", id)

	if drone.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Drone not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(drone)
}

func GetDroneByStation(c *fiber.Ctx) error {
	db := database.DB.Db

	id := c.Params("id")

	var drones []model.Drone

	db.Find(&drones, "station_id = ?", id)

	if len(drones) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Drone not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(drones)
}
