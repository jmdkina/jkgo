package jkbase

import (
	"html/template"
	"io"
)

type SimpleParse struct {
}

func (sp *SimpleParse) ParseString(out io.Writer, content string, data interface{}) error {
	t, err := template.New("simpleparse").Parse(content)
	if err != nil {
		return err
	}
	t.Execute(out, data)
	return nil
}

func (sp *SimpleParse) Parse(out io.Writer, file string, data interface{}) error {
	t, err := template.ParseFiles(file)
	if err != nil {
		return err
	}
	return t.Execute(out, data)
}

type WebResult struct {
	Desc string
}

func (wr *WebResult) NewWebResult(desc string) *WebResult {
	return &WebResult{
		Desc: desc,
	}
}

func (wr *WebResult) NewWebResultNoPermission() *WebResult {
	return &WebResult{
		Desc: "No Permission",
	}
}

func (w *WebResult) NewWebResultNotExist() *WebResult {
	return &WebResult{
		Desc: "Not exist",
	}
}
