package unsafe

import "unsafe"

func GetSliceLen[T any](s []T) int {
	return *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
}

func GetSliceCap[T any](s []T) int {
	return *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
}
