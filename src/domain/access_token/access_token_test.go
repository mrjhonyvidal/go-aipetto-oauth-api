package access_token

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	/**if expirationTime != 24 {
		t.Error("expiration should time be more than 24 hours")
	}**/
	assert.EqualValues(t, 24, expirationTime, "expiration should time be more than 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	/**if at.IsExpired() {
		t.Error("brand new access token should not be nil")
	}**/
	assert.False(t, at.IsExpired(), "brand new access token should not be nil")

	/**if at.AccessToken != "" {
		t.Error("new access token should not have defined access token id")
	}**/
	assert.EqualValues(t, "", at.AccessToken, "new access token should not have defined access token id")

	/**if at.UserId != 0 {
		t.Error("new access token should not have an associated user id")
	}**/
	assert.True(t, at.UserId == 0, "new access token should not have an associated user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	/**if !at.IsExpired() {
		t.Error("empty access token should be expired by default")
	}**/
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	/**if at.IsExpired() {
		t.Error("access token expiring three hours from now should NOT be expired")
	}**/
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token expiring three hours from now should NOT be expired")
}