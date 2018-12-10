package jkmisc

import (
	"testing"
)

func TestWetherQuery(t *testing.T) {
	w, _ := JKWetherNew("")
	_, err := w.Query("深圳")
	if err != nil {
		t.Fatal("error query ", err)
	}
}
