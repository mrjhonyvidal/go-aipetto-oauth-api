package db

import (
	"github.com/aipetto/go-aipetto-oauth-api/src/clients/cassandra"
	"github.com/aipetto/go-aipetto-oauth-api/src/domain/access_token"
	"github.com/aipetto/go-aipetto-oauth-api/src/utils/errors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken = "SELECT access_token, user_id, client_id, expires FROM access_token WHERE access_token=?;"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer session.Close()

	var result access_token.AccessToken
	if err := session.Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires,
		); err != nil {
			if err == gocql.ErrNotFound {
				return nil, errors.NewNotFoundError("no access token found with given id")
			}
			return nil, errors.NewInternalServerError(err.Error())
		}
	return &result, nil
}
