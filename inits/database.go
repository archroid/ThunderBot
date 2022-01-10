package inits

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDatabase() *mongo.Database {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	db, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logrus.Fatal(err)
	}

	// Check the connection
	err = db.Ping(context.TODO(), nil)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("Connected to MongoDB!")

	//return the whole database
	return db.Database("ElProfessorBot")
}
