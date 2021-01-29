package rest

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("About to start our test cases...")
	os.Exit(m.Run())
}

func TestLoginNotError(t *testing.T){
	repository := restUsersRepository{}
	user, err := repository.LoginUser("mrjhonyvidal@aipetto.com", "mypassword")

	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.EqualValues(t, 3, user.Id)
	assert.EqualValues(t, "Jhony", user.FirstName)
	assert.EqualValues(t, "Vidal", user.LastName)
	assert.EqualValues(t, "mrjhonyvidal@aipetto.com", user.Email)
}

func TestLoginUserTimeoutFromApi(t *testing.T){
	// TODO Understand httpmock - is not intercepting/mocking our real request
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// mock to add a new article
	httpmock.RegisterResponder("POST", "http://localhost:8081/users/login",
		func(req *http.Request) (*http.Response, error) {
				return httpmock.NewStringResponse(500, "internal server error"), nil
		},
	)
	/*
	TODO With Mock
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid rest client response when trying to login user", err.Message)
	*/
}

func TestLoginUserInvalidErrorInterface(t *testing.T){
	/**
	HttpMehotd: 	http.MethodPost
	URL: 			http://localhost:8081/users/login
	ReqBody: 		`{"email":"email@email.com,"password":"the-password"}`,
	RespHTTPCode:	http.StatusNotFound
	RespBody		`{"message":"invalid login credentials", "status":"404", "error": "not_found"}`

	repository := restUsersRepository{}
	user, err := repository.LoginUser("mrjhonyvidal@aipetto.com", "mypassword")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface when trying to login user", err.Message)
	*/
}

func TestLoginUserInvalidLoginCredentials(t *testing.T){
	/**
	HttpMehotd: 	http.MethodPost
	URL: 			http://localhost:8081/users/login
	ReqBody: 		`{"email":"email@email.com,"password":"the-password"}`,
	RespHTTPCode:	http.StatusNotFound
	RespBody		`{"message":"invalid login credentials", "status":404, "error": "not_found"}`

	repository := restUsersRepository{}
	user, err := repository.LoginUser("mrjhonyvidal@aipetto.com", "mypassword")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "invalid user credentials", err.Message)
	*/
}

func TestLoginInvalidUserJsonResponse(t *testing.T){
	/**
	HttpMehotd: 	http.MethodPost
	URL: 			http://localhost:8081/users/login
	ReqBody: 		`{"email":"email@email.com,"password":"the-password"}`,
	RespHTTPCode:	http.StatusInternalServerError
	RespBody		`{"message":"invalid login credentials", "status":404, "error": "not_found"}`

	repository := restUsersRepository{}
	user, err := repository.LoginUser("mrjhonyvidal@aipetto.com", "mypassword")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal login users response", err.Message)
	*/
}

