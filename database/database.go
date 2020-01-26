package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoCtx context.Context
var DBCli *mongo.Client
var err error

func ConnectToDB() (*mongo.Client, error) {
	fmt.Println("Connecting to MongoDB")

	//non-nil empty context
	mongoCtx := context.Background()

	// connecting to the mongodb
	DBCli, err = mongo.Connect(mongoCtx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}

	// check if the connection was successful by pining the MongoDB server
	err = DBCli.Ping(mongoCtx, nil)
	if err != nil {
		return nil, fmt.Errorf("could not connect to MongoDB: %v", err)
	}

	fmt.Println("Connection with MongoDB server has been established...")

	return DBCli, nil
}

func CloseConnection(dbClient *mongo.Client) error {
	fmt.Println("Closing MongoDB connection")
	err := dbClient.Disconnect(mongoCtx)
	if err != nil {
		return fmt.Errorf("could not close DB connection: %v", err)
	}

	return nil
}

//
func SpecifyCollection(dbClient *mongo.Client, dbName string, collectionName string) (*mongo.Collection, error) {
	// creating the collection blog on the mydb database
	blogDB := dbClient.Database(dbName).Collection(collectionName)

	return blogDB, nil
}
