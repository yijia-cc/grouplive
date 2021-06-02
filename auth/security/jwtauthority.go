package security

import (
	"encoding/json"

	"github.com/dgrijalva/jwt-go"
)

type JWTAuthority struct {
	signingKey []byte
}

func (j JWTAuthority) IssueToken(payload interface{}) (string, error) {
	payloadMap, err := toMap(payload)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(payloadMap))
	return token.SignedString(j.signingKey)
}

func (j JWTAuthority) GetPayload(token string, output *interface{}) {
	panic("Implement me!")
}

func toMap(input interface{}) (map[string]interface{}, error) {
	output := make(map[string]interface{})
	jsonBuf, _ := json.Marshal(input)
	err := json.Unmarshal(jsonBuf, &output)
	return output, err
}

func NewJWTAuthority(signingKey string) JWTAuthority {
	return JWTAuthority{signingKey: []byte(signingKey)}
}
