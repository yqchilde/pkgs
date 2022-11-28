package unsafe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSliceLenAndCap(t *testing.T) {
	s := make([]int, 3, 5)
	len := GetSliceLen(s)
	cap := GetSliceCap(s)
	assert.Equal(t, 3, len)
	assert.Equal(t, 5, cap)
}
