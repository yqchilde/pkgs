package utils

import (
	"bytes"
	"runtime"
	"strconv"
)

// GetGoroutineID generate a goroutine id
func GetGoroutineID() int {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, e := strconv.Atoi(string(b))
	if e != nil {
		return -1
	}

	return n
}
