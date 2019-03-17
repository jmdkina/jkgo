package simpleserver

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

func (sp *SimpleParse) Parse(out io.Writer, data interface{}, file ...string) error {
	t, err := template.ParseFiles(file...)
	if err != nil {
		return err
	}
	return t.Execute(out, data)
}
