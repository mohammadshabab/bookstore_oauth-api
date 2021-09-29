package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	// if expirationTime != 24 {
	// 	t.Error("expiration time should be 24 hours")
	// }
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken(1)
	assert.False(t, at.IsExpired(), "brand new access token should not be expired")
	assert.EqualValues(t, "", at.AccessToken, "new access token should not have defined access token id")
	assert.True(t, at.UserId == 0, "new access token should not have an associated user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "empty access token should be expired by default")
	// if !at.IsExpired() {
	// 	t.Error("empty access token should be expired by default")
	// }
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	// if at.IsExpired() {
	// 	t.Error("access token create three hours from now should NOT be expired")
	// }
	assert.False(t, at.IsExpired(), "access token create three hours from now should NOT be expired")
}
