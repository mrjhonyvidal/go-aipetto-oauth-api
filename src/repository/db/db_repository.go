package db

import (
	"errors"
	"github.com/aipetto/go-aipetto-oauth-api/src/clients/cassandra"
	"github.com/aipetto/go-aipetto-oauth-api/src/domain/access_token"
	"github.com/aipetto/go-aipetto-utils/src/rest_errors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken 					= "SELECT access_token, user_id, client_id, expires FROM access_token WHERE access_token=?;"
	queryCreateAccessToken 					= "INSERT INTO access_token(access_token, user_id, client_id, expires) VALUES(?,?,?,?);"
	queryUpdateAccessTokenExpirationTime	= "UPDATE access_token SET expires=? WHERE access_token=?;"

)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *rest_errors.RestErr)
	Create(token access_token.AccessToken) *rest_errors.RestErr
	UpdateExpirationTime(token access_token.AccessToken) *rest_errors.RestErr
}

type dbRepository struct {}

func (r *dbRepository) Create(at access_token.AccessToken) *rest_errors.RestErr {
	if err := cassandra.GetSession().Query(queryCreateAccessToken,
			at.AccessToken,
			at.UserId,
			at.ClientId,
			at.Expires).Exec(); err != nil {
			return rest_errors.NewInternalServerError(err.Error(), errors.New("Database Error"))
	}
	return nil
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *rest_errors.RestErr) {
	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires,
		); err != nil {
			if err == gocql.ErrNotFound {
				return nil, rest_errors.NewNotFoundError("no access token found with given id")
			}
			return nil, rest_errors.NewInternalServerError(err.Error(), errors.New("Access Token Invalid"))
		}
	return &result, nil
}

func (r *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *rest_errors.RestErr {
	if err := cassandra.GetSession().Query(queryUpdateAccessTokenExpirationTime,
		at.Expires,
		at.AccessToken,
	).Exec(); err != nil {
		return rest_errors.NewInternalServerError(err.Error(), errors.New("Database Error on Expiration Time Renew"))
	}
	return nil
}
