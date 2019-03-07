package jkitoy

import (
	"testing"
)

func TestItoY_AllDoItem(t *testing.T) {
	itoy, err := NewItoYItem("/Users/jmdvirus/Documents/1.jpg")
	if err != nil {
		t.Fatal(err)
	}
	err = itoy.DoSomething()
	if err != nil {
		t.Fatalf("error doing %v\n", err)
	}
}

func TestItoY_Travers(t *testing.T) {
	itoy, err := NewItoYItem("/Users/jmdvirus/Documents/item/1.bin")
	if err != nil {
		t.Fatal(err)
	}
	err = itoy.Travers()
	if err != nil {
		t.Fatal(err)
	}
}