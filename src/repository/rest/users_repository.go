package rest

import (
	"encoding/json"
	"fmt"
	"github.com/aipetto/go-aipetto-oauth-api/src/domain/users"
	"github.com/aipetto/go-aipetto-oauth-api/src/utils/errors"
	"github.com/go-resty/resty/v2"
	"time"
)

type RestUsersRepository interface{
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

func NewRepository() RestUsersRepository {
	return &restUsersRepository{}
}

type restUsersRepository struct{}

func (r *restUsersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {

	request := users.UserLoginRequest{
		Email: 		email,
		Password: 	password,
	}

	client := resty.New().SetHostURL("https://localhost:8081").SetTimeout(1 * time.Minute)

	resp, err := client.R().
					SetBody(request).Post("/users/login")

	if err != nil {
		return nil, errors.NewInternalServerError("invalid rest client response when trying to login user")
	}

	if resp.StatusCode() > 299 {
		var restErr errors.RestErr

		// TODO check if use Marshal/Unmarshal or NewDecoder/Decode
		err := json.Unmarshal(resp.Body(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user")
		}
		return nil, &restErr
	}

	var user users.User

	if err := json.Unmarshal(resp.Body(), &user); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal users response")
	}


	fmt.Print(resp)
	fmt.Print(err)

	return &user, nil
}
