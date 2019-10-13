package main

import (
	"reflect"
	"unsafe"
)

var pubBuf []byte
func ConvertToTitle(n int) string {
	// buf max size: 7, since 32*(log(2)/log(26)) < 7
	// if on 64bit system, 14 is enough.
	buf := make([]byte ,7)
	startIndex := 6
	for; startIndex >= 0 && n>0;startIndex--{
		n = n - 1
		m := n % 26
		n = n / 26
		buf[startIndex] = byte(m + 'A')
	}
	return string(buf[startIndex+1:])
}
func ConvertToTitleStringConcat(n int) string {
	// buf max size: 7, since 32*(log(2)/log(26)) < 7
	// if on 64bit system, 14 is enough.
	buf := ""
	startIndex := 6
	for; startIndex >= 0 && n>0;startIndex--{
		n = n - 1
		m := n % 26
		n = n / 26
		buf = buf + string(byte(m + 'A'))
	}
	return buf
}
func ConvertToTitlePubBuf(n int) string {
	// buf max size: 7, since 32*(log(2)/log(26)) < 7
	// if on 64bit system, 14 is enough.
	if pubBuf == nil {
		pubBuf = make([]byte ,7)
	}
	startIndex := 6
	for; startIndex >= 0 && n>0;startIndex--{
		n = n - 1
		m := n % 26
		n = n / 26
		pubBuf[startIndex] = byte(m + 'A')
	}
	return *(*string)(unsafe.Pointer(&reflect.StringHeader{
		Data: ((*reflect.SliceHeader)(unsafe.Pointer(&pubBuf))).Data+uintptr(startIndex + 1),
		Len:  6 - startIndex,
	}))
}
func main() {

}