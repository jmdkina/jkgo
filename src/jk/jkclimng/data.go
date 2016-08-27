package jkclimng

import "encoding/json"

// Data define

// 1. GetItem
// { "Header":{ "ID":"1234", "Cmd":"GetItem"}, "Body":{ "Count": 1}  }
// { "Header":{ "ID":"1234", "Cmd":"GetItem"}, "Body":{ "Result":"success", "Count":1, "Data":[{"Content":"xxxx", "Abstract":"iii" }] } }

type AIHeader struct {
	ID  string
	Cmd string
}

type AIGeneral struct {
	Header AIHeader
	Body   interface{}
}

func (gen *AIGeneral) Parse(data []byte) error {
	err := json.Unmarshal(data, gen)
	return err
}

type AIGetItemRequestBody struct {
	Count int
}

type AIGetItemRequest struct {
	Header AIHeader
	Body   AIGetItemRequestBody
}

func (gi *AIGetItemRequest) Parse(data []byte) error {
	err := json.Unmarshal(data, gi)
	return err
}

type AIGetItemResponseBody struct {
	Result string
	Count  int
	Data   []MRYZItem
}

type AIGetItemResponse struct {
	Header AIHeader
	Body   AIGetItemResponseBody
}

func (gi *AIGetItemResponse) Parse(data []byte) error {
	err := json.Unmarshal(data, gi)
	return err
}

type MRYZItem struct {
	Content    string
	Author     string
	Abstract   string
	CreateTime int64
	UpdateTime int64
	Image      string
}
