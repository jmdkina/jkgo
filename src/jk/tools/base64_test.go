package tools

import (
	"testing"
)

func TestDecode(t *testing.T) {
    c := NewEncoding()
    old := "isl sfa\nkisa"
    vv := c.Encode([]byte(old))
    t.Log(vv)
    v, _ := c.Decode("5LiK6K+B5oyH5")
    t.Log(string(v))
    t.Fatal("test")
}