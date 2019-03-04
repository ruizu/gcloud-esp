package esp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseUserInfo(t *testing.T) {
	var testCases = []struct {
		key      string
		expected UserInfo
	}{
		{
			key: "eyJpc3N1ZXIiOiJhbmRyaXMubWUiLCJpZCI6InJ1aXp1IiwiZW1haWwiOiJsb3Vpc0BhbmRyaXMubWUifQ==",
			expected: UserInfo{
				Issuer: "andris.me",
				ID:     "ruizu",
				Email:  "louis@andris.me",
			},
		},
		{
			key: "eyJpZCI6InJ1aXp1IiwiZW1haWwiOiJsb3Vpc0BhbmRyaXMubWUifQ==",
			expected: UserInfo{
				ID:    "ruizu",
				Email: "louis@andris.me",
			},
		},
		{
			key: "eyJpc3N1ZXIiOiJhbmRyaXMubWUiLCJlbWFpbCI6ImxvdWlzQGFuZHJpcy5tZSJ9",
			expected: UserInfo{
				Issuer: "andris.me",
				Email:  "louis@andris.me",
			},
		},
		{
			key: "eyJpc3N1ZXIiOiJhbmRyaXMubWUiLCJpZCI6InJ1aXp1In0=",
			expected: UserInfo{
				Issuer: "andris.me",
				ID:     "ruizu",
			},
		},
		{
			key:      "e30=",
			expected: UserInfo{},
		},
	}

	var res UserInfo
	var err error
	for _, v := range testCases {
		res, err = parseUserInfo(v.key)
		assert.Nil(t, err)
		assert.Equal(t, res, v.expected)
	}
}
