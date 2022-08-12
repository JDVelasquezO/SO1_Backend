package Controller

import (
	"backPractica1_SO1/Models"
	"backPractica1_SO1/instance"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetVehicles(c *fiber.Ctx) error {
	query := bson.D{{}}

	cursor, err := instance.Mg.Db.Collection("employees").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	var vehicles = make([]Models.Employee, 0)

	if err := cursor.All(c.Context(), &vehicles); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(vehicles)
}
