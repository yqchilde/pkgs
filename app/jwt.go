package app

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	ErrMissingAuthorization = errors.New("the authorization information is incorrect")
)

type Payload struct {
	UserID string
}

// Sign signs the payload with the specified secret
// iss: （Issuer）Issuer
// iat: （Issued At）Issuance time, expressed in Unix timestamp
// exp: （Expiration Time）Expiration time, expressed in Unix timestamp
// aud: （Audience）The party receiving the JWT
// sub: （Subject）The subject of the JWT
// nbf: （Not Before）Not earlier than this time
// jti: （JWT ID）Unique ID used to identify JWT
func Sign(payload map[string]interface{}, secret string, timeout int64) (tokenStr string, err error) {
	now := time.Now().Unix()
	claims := make(jwt.MapClaims)
	claims["nbf"] = now
	claims["iat"] = now
	claims["exp"] = now + timeout

	for k, v := range payload {
		claims[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err = token.SignedString([]byte(secret))

	return
}

func ParseRequest(token, secret string) (*Payload, error) {
	if token == "" {
		return &Payload{}, ErrMissingAuthorization
	}

	var t string
	_, err := fmt.Sscanf(token, "Bearer %s", &t)
	if err != nil {
		fmt.Printf("fmt.Sscanf err: %v", err)
	}

	return Parse(t, secret)
}

func Parse(tokenStr string, secret string) (*Payload, error) {
	token, err := jwt.Parse(tokenStr, secretFunc(secret))
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		payloads := &Payload{
			UserID: claims["user_id"].(string),
		}
		return payloads, nil
	}
	return nil, err
}

func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}
