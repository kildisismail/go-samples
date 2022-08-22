package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func Reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func Reverse2(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev := Reverse(input)
	doubleRev := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)

	rev, revErr := Reverse2(input)
	doubleRev, doubleRevErr := Reverse2(rev)
	fmt.Printf("\noriginal2: %q\n", input)
	fmt.Printf("reversed2: %q, err: %v\n", rev, revErr)
	fmt.Printf("reversed again2: %q, err: %v\n", doubleRev, doubleRevErr)
}
