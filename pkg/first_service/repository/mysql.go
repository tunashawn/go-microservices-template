package repository

import (
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	db2 "go-microservices-template/internal/db"
)

type SqlRepository interface {
}

type SqlRepositoryImpl struct {
	db *bun.DB
}

func NewSqlRepository() (SqlRepository, error) {
	db, err := db2.NewMySQLConnection()
	if err != nil {
		return nil, errors.Wrap(err, "could not create mysql connection")
	}
	return SqlRepositoryImpl{db: db}, nil
}
