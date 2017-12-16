package jkbase

import (
	"encoding/base64"
)

type Base64 struct {
	coder *base64.Encoding
}

func NewEncoding() *Base64 {
	base64Table := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	coder := base64.NewEncoding(base64Table)
	return &Base64{coder: coder}
}

func (b *Base64) Encode(data []byte) string {
	return b.coder.EncodeToString(data)
}

func (b *Base64) Decode(data string) ([]byte, error) {
	return b.coder.DecodeString(data)
}
