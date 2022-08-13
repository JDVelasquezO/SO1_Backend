package Controller

import (
	"backPractica1_SO1/Models"
	"backPractica1_SO1/instance"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func PostRecord(c fiber.Ctx) error {
	collectionRecord := instance.Mg.Db.Collection("record")

	record := new(Models.Record)
	record.ID = ""
	record.Func = "Register"
	currentTime := time.Now()
	record.Time = currentTime.Format("2006-01-02 15:04:05")

	insertRes, err := collectionRecord.InsertOne(c.Context(), record)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	filter := bson.D{{
		Key:   "_id",
		Value: insertRes.InsertedID,
	}}

	createdRecord := collectionRecord.FindOne(c.Context(), filter)

	createdRecordDB := &Models.Record{}
	err = createdRecord.Decode(createdRecordDB)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(createdRecordDB)
}
