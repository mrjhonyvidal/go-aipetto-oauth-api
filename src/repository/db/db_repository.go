package db

import (
	"github.com/aipetto/go-aipetto-oauth-api/src/domain/access_token"
	"github.com/aipetto/go-aipetto-oauth-api/src/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr){
	return nil, nil
}
