package jkclimng

import "testing"

func TestAIGeneral_Parse(t *testing.T) {
	str := "{ \"Header\":{ \"ID\":\"1234\", \"Cmd\":\"GetItem\"}, \"Body\":{ \"Count\": 1}  }"
	gen := AIGeneral{}
	err := gen.Parse([]byte(str))
	if err != nil {
		t.Fatal("error parse ", err)
	}
	if gen.Header.ID != "1234" {
		t.Fatal("ID need 1234, but real is ", gen.Header.ID)
	}
	if gen.Header.Cmd != "GetItem" {
		t.Fatal("Cmd need GetItem, but real is ", gen.Header.Cmd)
	}
}

func TestAIGetItemRequest_Parse(t *testing.T) {
	str := "{ \"Header\":{ \"ID\":\"1234\", \"Cmd\":\"GetItem\"}, \"Body\":{ \"Count\": 1}  }"
	req := AIGetItemRequest{}
	err := req.Parse([]byte(str))
	if err != nil {
		t.Fatal("error parse ", err)
	}
	if req.Header.ID != "1234" {
		t.Fatal("ID need 1234, but real is ", req.Header.ID)
	}
	if req.Header.Cmd != "GetItem" {
		t.Fatal("Cmd need GetItem, but real is ", req.Header.Cmd)
	}
	if req.Body.Count != 1 {
		t.Fatal("Count need 1, but real is ", req.Body.Count)
	}
}

func TestAIGetItemResponse_Parse(t *testing.T) {
	str := "{ \"Header\":{ \"ID\":\"1234\", \"Cmd\":\"GetItem\"}, \"Body\":{ \"Count\": 1, \"Result\": \"success\", \"Data\":[{ \"Content\":\"doning\", \"Author\":\"jmd\" }]}  }"
	res := AIGetItemResponse{}
	err := res.Parse([]byte(str))
	if err != nil {
		t.Fatal("error parse ", err)
	}
	if res.Header.ID != "1234" {
		t.Fatal("ID need 1234, but real is ", res.Header.ID)
	}
	if res.Header.Cmd != "GetItem" {
		t.Fatal("Cmd need GetItem, but real is ", res.Header.Cmd)
	}
	if res.Body.Count != 1 {
		t.Fatal("Count need 1, but real is ", res.Body.Count)
	}
	items := res.Body.Data
	if len(items) != 1 {
		t.Fatal("Len items need 1, but real is ", len(items))
	}
	if items[0].Content != "doning" {
		t.Fatal("Content need doning, but real is ", items[0].Content)
	}
}
