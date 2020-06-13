package db

import (
	"github.com/edilbertloquine/go-microservices/oauth-api/src/clients/cassandra"
	"github.com/edilbertloquine/go-microservices/oauth-api/src/domain/access_token"
	"github.com/edilbertloquine/go-microservices/oauth-api/src/utils/errors"
)

// DbRepository -
type DbRepository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct{}

// NewRepository -
func NewRepository() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetByID(id string) (*access_token.AccessToken, *errors.RestErr) {
	_, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	return nil, errors.NewInternalServerError("database connection not established")
}
