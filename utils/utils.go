package utils

import (
	"bytes"
	"math/rand"
	"runtime"
	"strconv"
	"time"
)

func RandomStr(n int) string {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	const pattern = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz"

	salt := make([]byte, 0, n)
	l := len(pattern)

	for i := 0; i < n; i++ {
		p := r.Intn(l)
		salt = append(salt, pattern[p])
	}

	return string(salt)
}

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
