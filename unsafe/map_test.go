package unsafe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMapLen(t *testing.T) {
	m := make(map[string]interface{}, 0)
	m["name"] = "yqchilde"
	m["age"] = 18
	len := GetMapLen(m)
	assert.Equal(t, 2, len)
}
