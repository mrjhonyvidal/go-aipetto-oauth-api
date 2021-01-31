package rest

import (
	"encoding/json"
	"errors"
	"github.com/aipetto/go-aipetto-oauth-api/src/domain/users"
	"github.com/aipetto/go-aipetto-utils/src/rest_errors"
	"github.com/go-resty/resty/v2"
	"time"
)

type RestUsersRepository interface{
	LoginUser(string, string) (*users.User, *rest_errors.RestErr)
}

func NewRestUserRepository() RestUsersRepository {
	return &restUsersRepository{}
}

type restUsersRepository struct{}

func (r *restUsersRepository) LoginUser(email string, password string) (*users.User, *rest_errors.RestErr) {

	request := users.UserLoginRequest{
		Email: 		email,
		Password: 	password,
	}

	client := resty.New().SetHostURL("http://localhost:8081").SetTimeout(1 * time.Minute)
	resp, err := client.R().
				SetBody(request).Post("/users/login")

	if err != nil {
		return nil, rest_errors.NewInternalServerError("invalid rest client response when trying to login user", errors.New("Rest Server Error"))
	}

	if resp.StatusCode() > 299 {
		var restErr rest_errors.RestErr

		err := json.Unmarshal(resp.Body(), &restErr)
		if err != nil {
			return nil, rest_errors.NewInternalServerError("invalid error interface when trying to login user", errors.New("JSON Decode/Unmarshall login error"))
		}
		return nil, &restErr
	}

	var user users.User

	if err := json.Unmarshal(resp.Body(), &user); err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to unmarshal login users response", errors.New("JSON Decode/Unmarshall login response error "))
	}

	return &user, nil
}
