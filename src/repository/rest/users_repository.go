package rest

import (
	"fmt"
	"github.com/aipetto/go-aipetto-oauth-api/src/domain/users"
	"github.com/aipetto/go-aipetto-oauth-api/src/utils/errors"
	"github.com/go-resty/resty/v2"

)

type RestUsersRepository interface{
	LoginUser(string, string) (*users.User, errors.RestErr)
}

func NewRepository() RestUsersRepository {
	return &restUsersRepository{}
}

type restUsersRepository struct{}

func (r *restUsersRepository) LoginUser(email string, password string) (*users.User, errors.RestErr) {

	request := users.UserLoginRequest{
		Email: 		email,
		Password: 	password,
	}
	client := resty.New()

	resp, err := client.R().
					SetBody(request).Post("/users/login")

	fmt.Print(resp)
	fmt.Print(err)

	return nil, errors.RestErr{}
}
