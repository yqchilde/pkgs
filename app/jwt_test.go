package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSign(t *testing.T) {
	jwtPayload := map[string]interface{}{"user_id": "yqchilde", "email": "yqchilde@gmail.com"}
	jwtSecret := "jwt_secret"
	jwtExpire := int64(3600)
	token, err := Sign(jwtPayload, jwtSecret, jwtExpire)
	if err != nil {
		t.Error(err)
	}

	parsePayload, err := ParseRequest("Bearer "+token, jwtSecret)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, jwtPayload["user_id"], parsePayload.UserID)
}
