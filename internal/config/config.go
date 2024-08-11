package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type MongoDBConfig struct {
	ConnectionURI string `envconfig:"MONGODB_CONNECTION_URI" required:"true"`
	DatabaseName  string `envconfig:"MONGODB_DATABASE_NAME" required:"true"`
}

type MySQLConfig struct {
	ConnectionString string `envconfig:"MYSQL_CONNECTION_STRING" required:"true"`
}

func GetConfig(cfg any) error {
	err := envconfig.Process("", cfg)
	if err != nil {
		return errors.Wrap(err, "could not read auth_service config")
	}

	return nil
}
