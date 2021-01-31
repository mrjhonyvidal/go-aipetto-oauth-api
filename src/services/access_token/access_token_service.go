package access_token

import (
	"github.com/aipetto/go-aipetto-oauth-api/src/domain/access_token"
	"github.com/aipetto/go-aipetto-oauth-api/src/repository/db"
	"github.com/aipetto/go-aipetto-oauth-api/src/repository/rest"
	"github.com/aipetto/go-aipetto-utils/src/rest_errors"
	"strings"
)

type Service interface {
	GetById(string) (*access_token.AccessToken, *rest_errors.RestErr)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, *rest_errors.RestErr)
	UpdateExpirationTime(access_token.AccessToken) *rest_errors.RestErr
}

type service struct {
	dbRepo db.DbRepository
	restRepo rest.RestUsersRepository
}

func NewService(dbRepository db.DbRepository, restRepository rest.RestUsersRepository) Service {
	return &service{
		dbRepo: dbRepository,
		restRepo: restRepository,
	}
}

func (s *service) GetById(accessTokenId string) (*access_token.AccessToken, *rest_errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, rest_errors.NewBadRequestError("invalid access token id")
	}

	accessToken, err := s.dbRepo.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, err
}

func (s *service) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, *rest_errors.RestErr) {

	if err := request.ValidateAccessToken(); err != nil{
		return nil, err
	}

	// TODO: Suport both grant types: client_credentials and password

	user, err := s.restRepo.LoginUser(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	// Authenticate the user against the Users API:
	at := access_token.GetNewAccessToken(user.Id)
	at.Generate()

	// Save the new access token in Cassandra
	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}

	return &at, nil
}

func (s *service) UpdateExpirationTime(at access_token.AccessToken) *rest_errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.dbRepo.UpdateExpirationTime(at)
}