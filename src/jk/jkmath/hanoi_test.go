package jkmath

import (
	"testing"
)

func TestHanoi(t *testing.T) {
	n := 8
	h := NewHanoi(n)
	h.Do()
	h.Print()
}
