package services

import (
	"github.com/pkg/errors"
	"go-microservices-template/pkg/service_A/repository"
)

type Service interface {
	Ping() (string, error)
}

type ServiceImpl struct {
	mongo repository.MongoRepository
	sql   repository.SqlRepository
}

func (s *ServiceImpl) Ping() (string, error) {
	return "oke", nil
}

func NewService() (Service, error) {
	mongo, err := repository.NewMongoRepository()
	if err != nil {
		return nil, errors.Wrap(err, "could not create new mongo repository")
	}

	sql, err := repository.NewSqlRepository()
	if err != nil {
		return nil, errors.Wrap(err, "could not create new sql repository")
	}

	return &ServiceImpl{
		mongo: mongo,
		sql:   sql,
	}, nil
}
