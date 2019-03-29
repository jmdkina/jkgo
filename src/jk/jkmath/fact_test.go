package jkmath

import (
	"testing"
)

func TestFact(t *testing.T) {
	n := 8
	v := Fact(n)
	t.Log(v)
}
