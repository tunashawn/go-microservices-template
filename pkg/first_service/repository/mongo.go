package repository

import (
	"github.com/pkg/errors"
	db2 "go-microservices-template/internal/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository interface {
}

type MongoRepositoryImpl struct {
	db *mongo.Database
}

func NewMongoRepository() (MongoRepository, error) {
	db, err := db2.NewMongoDatabase()
	if err != nil {
		return nil, errors.Wrap(err, "could not create new mongodb connection")
	}
	return MongoRepositoryImpl{db: db}, nil
}
