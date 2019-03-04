package esp

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrEmptyUserInfo = errors.New("no user info header present in the request")

	errDecodeUserInfo = errors.New("unable to decode user info")
)

// UserInfo object contains information of the current logged in user.
type UserInfo struct {
	Issuer string `json:"issuer"`
	ID     string `json:"id"`
	Email  string `json:"email"`
}

// GetUserInfo returns the UserInfo sent by Google Cloud ESP.
func GetUserInfo(r *http.Request) (u UserInfo, err error) {
	encodedUserInfo := r.Header.Get("X-Endpoint-API-UserInfo")
	if len(encodedUserInfo) == 0 {
		return u, ErrEmptyUserInfo
	}

	return parseUserInfo(encodedUserInfo)
}

func parseUserInfo(encodedUserInfo string) (u UserInfo, err error) {
	var decodedUserInfo []byte
	decodedUserInfo, err = base64.URLEncoding.DecodeString(encodedUserInfo)
	if err != nil {
		return u, errDecodeUserInfo
	}

	err = json.Unmarshal(decodedUserInfo, &u)
	if err != nil {
		return u, errDecodeUserInfo
	}

	return
}
