package Database

import (
	"backPractica1_SO1/instance"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

const dbName = "DB"
const mongoURI = "mongodb://localhost:27017/" + dbName

func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return err
	}

	instance.Mg = instance.MongoInstance(MongoInstance{
		Client: client,
		Db:     db,
	})

	return nil
}
