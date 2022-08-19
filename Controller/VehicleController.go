package Controller

import (
	"backPractica1_SO1/Models"
	"backPractica1_SO1/instance"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetVehicles(c *fiber.Ctx) error {
	query := bson.D{{}}

	cursor, err := instance.Mg.Db.Collection("vehicles").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	var vehicles = make([]Models.Vehicle, 0)

	if err := cursor.All(c.Context(), &vehicles); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(vehicles)
}

func PostVehicles(c *fiber.Ctx) error {
	// Keep Vehicles in DB
	collection := instance.Mg.Db.Collection("vehicles")
	vehicle := new(Models.Vehicle)

	if err := c.BodyParser(vehicle); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	insertRes, err := collection.InsertOne(c.Context(), vehicle)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	filter := bson.D{{
		Key:   "_id",
		Value: insertRes.InsertedID,
	}}

	createdRecord := collection.FindOne(c.Context(), filter)

	createdVehicle := &Models.Vehicle{}
	err = createdRecord.Decode(createdVehicle)
	if err != nil {
		return err
	}

	// Keep Date and Time
	err = PostRecord(*c, "register")
	if err != nil {
		return err
	}

	return c.Status(201).JSON(createdVehicle)
}

func PutVehicles(c *fiber.Ctx) error {
	idVehicle := c.Params("id")

	vehicle := new(Models.Vehicle)

	if err := c.BodyParser(vehicle); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	query := bson.D{{Key: "_id", Value: idVehicle}}
	update := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "_id", Value: vehicle.ID},
				{Key: "model", Value: vehicle.Model},
				{Key: "series", Value: vehicle.Series},
				{Key: "mark", Value: vehicle.Mark},
				{Key: "color", Value: vehicle.Color},
			},
		},
	}

	err := instance.Mg.Db.Collection("vehicles").FindOneAndUpdate(c.Context(),
		query, update).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.SendStatus(400)
		}
		return c.SendStatus(500)
	}
	vehicle.ID = idVehicle

	// Keep Date and Time
	err = PostRecord(*c, "update")
	if err != nil {
		return err
	}

	return c.Status(200).JSON(vehicle)
}

func DelVehicles(c *fiber.Ctx) error {
	vehicleId := c.Params("id")

	query := bson.D{
		{
			Key:   "_id",
			Value: vehicleId,
		},
	}
	result, err := instance.Mg.Db.Collection("vehicles").DeleteOne(c.Context(), &query)

	if err != nil {
		return c.SendStatus(500)
	}

	if result.DeletedCount < 1 {
		return c.SendStatus(404)
	}

	// Keep Date and Time
	err = PostRecord(*c, "delete")
	if err != nil {
		return err
	}

	return c.Status(200).JSON("record deleted")
}
