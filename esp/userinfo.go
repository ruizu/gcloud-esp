package esp

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
)

// UserData contains decoded information from header sent by Cloud Endpoints ESP.
type UserData struct {
	Issuer string `json:"issuer"`
	ID     string `json:"id"`
	Email  string `json:"email"`
}

var errNoUserInfo = errors.New("esp: no user info header present in the request")

const headerField = "X-Endpoint-API-UserInfo"

// User decodes the user info header sent by Cloud Endpoints ESP.
func User(r *http.Request) (*UserData, error) {
	encodedData := r.Header.Get(headerField)
	if encodedData == "" {
		return nil, errNoUserInfo
	}

	decodedData, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		return nil, err
	}

	var user UserData
	if err := json.Unmarshal(decodedData, &user); err != nil {
		return nil, err
	}

	return &user, err
}
