package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncryptedPassword(t *testing.T) {
	password := "123456qweASD@!"
	hash, err := HashAndSalt(password)
	if err != nil {
		t.Error(err)
	}
	t.Logf("hash: %s", hash)
	assert.Equal(t, true, ComparePasswords(hash, password))
}
