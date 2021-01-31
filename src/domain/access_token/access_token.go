package access_token

import (
	"fmt"
	"github.com/aipetto/go-aipetto-oauth-api/src/utils/crypto"
	"github.com/aipetto/go-aipetto-utils/src/rest_errors"
	"strings"
	"time"
)

const (
	expirationTime = 24
	grantTypePassword = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType		string	`json:"grant_type"`
	Scope			string	`json:"scope"`

	// Used for password grant_type
	Username		string	`json:"username"`
	Password		string	`json:"password"`

	// User for client_credentials grant type
	ClientId		string	`json:"client_id"`
	ClientSecret	string	`json:"client_secret"`
}

func (at *AccessTokenRequest) ValidateAccessToken() *rest_errors.RestErr{
	switch at.GrantType{
	case grantTypePassword:
		break
	case grantTypeClientCredentials:
		break
	default:
		return rest_errors.NewBadRequestError("invalid grant_type parameter")
	}

	// Validate parameters for each grant_type
	return nil
}

type AccessToken struct {
	AccessToken string 	`json:"access_token"`
	UserId 		int64	`json:"user_id"`
	ClientId	int64	`json:"client_id,omitempty"`
	Expires		int64	`json:expires`
}

func (at *AccessToken) Validate() *rest_errors.RestErr{
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return rest_errors.NewBadRequestError("invalid access token id")
	}
	if at.UserId <= 0 {
		return rest_errors.NewBadRequestError("invalid user id")
	}
	if at.ClientId<= 0 {
		return rest_errors.NewBadRequestError("client id")
	}
	if at.Expires <= 0 {
		return rest_errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}

func GetNewAccessToken(userId int64) AccessToken{
	return AccessToken{
		UserId: userId,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto.GetTokenMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}