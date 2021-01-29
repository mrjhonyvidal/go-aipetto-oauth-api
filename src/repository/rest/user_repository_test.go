package rest

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"
	"github.com/jarcoal/httpmock"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T){
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://aipetto.com",
		func(req *http.Request) (*http.Response, error){
			user := make(map[string]interface{})
			if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
				return httpmock.NewStringResponse(404, ""), nil
			}
			return nil, nil
		})
}

func TestLoginUserInvalidErrorInterface(t *testing.T){}

func TestLoginUserInvalidLoginCredentials(t *testing.T){}

func TestLoginInvalidUserJsonResponse(t *testing.T){}

func TestLoginUserNoError(t *testing.T){}