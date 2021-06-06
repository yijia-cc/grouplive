package security

import (
	"encoding/json"
	"errors"

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

func (j JWTAuthority) GetPayload(tokenStr string, output interface{}) error {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return j.signingKey, nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("token is invalid")
	}

	var claims jwt.MapClaims
	var ok bool
	if claims, ok = token.Claims.(jwt.MapClaims); !ok {
		return errors.New("token payload is not map")
	}

	buf, err := json.Marshal(map[string]interface{}(claims))
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, output)
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
