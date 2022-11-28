package unsafe

import "unsafe"

func GetMapLen[K comparable, V any](m map[K]V) int {
	return **(**int)(unsafe.Pointer(&m))
}
