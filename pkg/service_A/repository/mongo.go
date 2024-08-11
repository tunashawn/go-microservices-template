package repository

import (
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
		return nil, err
	}
	return MongoRepositoryImpl{db: db}, nil
}
