package jkmisc

import (
	"testing"
)

func TestWetherQuery(t *testing.T) {
	w, _ := JKWetherNew("")
	r, err := w.Query("深圳")
	if err != nil {
		t.Fatal("error query ", err)
	}
	if r.Error_code != 0 {
		t.Errorf("error code should 0 but, give %d\n", r.Error_code)
	}
	if r.Result.Today.City != "深圳" {
		t.Fatalf("city should 深圳, but give %s\n", r.Result.Today.City)
	}
}
