package db

import (
	"github.com/mohammadshabab/bookstore_oauth-api/src/domain/access_token"
	"github.com/mohammadshabab/bookstore_oauth-api/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func (r *dbRepository) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("database connection not implemented yet")

}
