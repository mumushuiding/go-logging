package test

import (
	"fmt"
	"testing"
)

func TestTest(t *testing.T) {
	var a = 1
	var b = &a
	var c = *b
	*b = 3
	c = 2
	*b = c
	fmt.Printf("a=%d,b=%d,c=%d", a, *b, c)
}
