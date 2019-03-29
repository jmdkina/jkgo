package jklog

import "testing"

func TestNewLogWrap(t *testing.T) {
	lw := LW()
	err := lw.Init("/tmp/log.test")
	if err != nil {
		t.Fatal("Error init logs ", err)
	}

	t.Fatal("Give print")
}