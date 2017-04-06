package esp

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://example.com", nil)

	var user, expect *UserData
	var err error

	req.Header.Set("X-Endpoint-API-UserInfo", "eyJpc3N1ZXIiOiJhbmRyaXMubWUiLCJpZCI6InJ1aXp1IiwiZW1haWwiOiJsb3Vpc0BhbmRyaXMubWUifQ==")
	expect = &UserData{
		Issuer: "andris.me",
		ID: "ruizu",
		Email: "louis@andris.me",
	}
	user, err = User(req)
	assert.Nil(t, err)
	assert.Equal(t, user, expect)

	req.Header.Set("X-Endpoint-API-UserInfo", "eyJpZCI6InJ1aXp1IiwiZW1haWwiOiJsb3Vpc0BhbmRyaXMubWUifQ==")
	expect = &UserData{
		ID: "ruizu",
		Email: "louis@andris.me",
	}
	user, err = User(req)
	assert.Nil(t, err)
	assert.Equal(t, user, expect)

	req.Header.Set("X-Endpoint-API-UserInfo", "eyJpc3N1ZXIiOiJhbmRyaXMubWUiLCJlbWFpbCI6ImxvdWlzQGFuZHJpcy5tZSJ9")
	expect = &UserData{
		Issuer: "andris.me",
		Email: "louis@andris.me",
	}
	user, err = User(req)
	assert.Nil(t, err)
	assert.Equal(t, user, expect)

	req.Header.Set("X-Endpoint-API-UserInfo", "eyJpc3N1ZXIiOiJhbmRyaXMubWUiLCJpZCI6InJ1aXp1In0=")
	expect = &UserData{
		Issuer: "andris.me",
		ID: "ruizu",
	}
	user, err = User(req)
	assert.Nil(t, err)
	assert.Equal(t, user, expect)

	req.Header.Set("X-Endpoint-API-UserInfo", "eyJpc3N1ZXIiOiJhbmRyaXMubWUifQ==")
	expect = &UserData{
		Issuer: "andris.me",
	}
	user, err = User(req)
	assert.Nil(t, err)
	assert.Equal(t, user, expect)

	req.Header.Set("X-Endpoint-API-UserInfo", "eyJpZCI6InJ1aXp1In0=")
	expect = &UserData{
		ID: "ruizu",
	}
	user, err = User(req)
	assert.Nil(t, err)
	assert.Equal(t, user, expect)

	req.Header.Set("X-Endpoint-API-UserInfo", "eyJlbWFpbCI6ImxvdWlzQGFuZHJpcy5tZSJ9")
	expect = &UserData{
		Email: "louis@andris.me",
	}
	user, err = User(req)
	assert.Nil(t, err)
	assert.Equal(t, user, expect)

	req.Header.Set("X-Endpoint-API-UserInfo", "e30=")
	expect = &UserData{}
	user, err = User(req)
	assert.Nil(t, err)
	assert.Equal(t, user, expect)

	req.Header.Set("X-Endpoint-API-UserInfo", "eyJpc3N1ZXIiOjEsImlkIjoicnVpenUiLCJlbWFpbCI6ImxvdWlzQGFuZHJpcy5tZSJ9")
	user, err = User(req)
	assert.NotNil(t, err)

	req.Header.Set("X-Endpoint-API-UserInfo", "eyJpc3N1ZXIiOiJhbmRyaXMubWUiLCJpZCI6MiwiZW1haWwiOiJsb3Vpc0BhbmRyaXMubWUifQ==")
	user, err = User(req)
	assert.NotNil(t, err)

	req.Header.Set("X-Endpoint-API-UserInfo", "eyJpc3N1ZXIiOiJhbmRyaXMubWUiLCJpZCI6InJ1aXp1IiwiZW1haWwiOjF9")
	user, err = User(req)
	assert.NotNil(t, err)

	req.Header.Set("X-Endpoint-API-UserInfo", "eyJpc3N1ZXIiOiJhbmRyaXMubWUiLCJpZCI6InJ1aXp1IiwiZW1haWwiOiJsb3Vpc0BhbmRyaXMubWUi")
	user, err = User(req)
	assert.NotNil(t, err)
}
