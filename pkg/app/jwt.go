package app

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/yqchilde/gint/pkg/config"
)

var (
	ErrMissingHeader = errors.New("the length of the `Authorization` header is zero")
)

type Payload struct {
	UserID      string
	AuthorityID string
}

// secretFunc validates the secret format.
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// Make sure the `alg` is what we except.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}

// Parse validates the token with the specified secret,
// and returns the payloads if the token was valid.
func Parse(tokenString string, secret string) (*Payload, error) {
	// Parse Token
	token, err := jwt.Parse(tokenString, secretFunc(secret))
	if err != nil {
		return nil, err
	}

	// Read the token if it's valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		payloads := &Payload{
			UserID:      claims["user_id"].(string),
			AuthorityID: claims["authority_id"].(string),
		}
		return payloads, nil
	}

	// Other errors
	return nil, err
}

// ParseRequests gets the token from the header and
// pass it to the Parse function to parses the token.
func ParseRequests(c *gin.Context) (*Payload, error) {
	authorization := c.Request.Header.Get("Authorization")

	// Load the jwt secret from config
	secret := config.Cfg.Server.JwtSecret

	if len(authorization) == 0 {
		return &Payload{}, ErrMissingHeader
	}

	var t string
	// Parse the authorization to get the token part
	_, err := fmt.Sscanf(authorization, "Bearer %s", &t)
	if err != nil {
		fmt.Printf("fmt.Sscanf err: %+v", err)
	}

	return Parse(t, secret)
}

func Sign(ctx context.Context, payload map[string]interface{}, secret string, timeout int64) (tokenString string, err error) {
	now := time.Now().Unix()
	claims := make(jwt.MapClaims)
	claims["nbf"] = now
	claims["iat"] = now
	claims["exp"] = now + timeout

	for k, v := range payload {
		claims[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the to ken with the specified secret.
	tokenString, err = token.SignedString([]byte(secret))

	return
}
