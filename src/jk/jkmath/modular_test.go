package jkmath

import (
"testing"
)

func TestBaseExpansion(t *testing.T) {
	n := 123
	b := 2
	m := New()
	m.BaseExpansion(n, b)
}