package jkeasycrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	// "encoding/hex"
	// "fmt"
	"io"
	"jk/jklog"
)

func JKAESEncrypt(key, text []byte) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		return ""
	}
	ciphertext := make([]byte, len(text)+aes.BlockSize)
	iv := ciphertext[:aes.BlockSize]

	// iv := []byte{
	// 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06,
	// 0x07, 0x08, 0x09, 0x0A, 0x0B,
	// 0x0C, 0x0D, 0x0E, 0x0F,
	// }

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return ""
	}
	// jklog.L().Infof("cipher text: %v\n", iv)

	cfb := cipher.NewCFBEncrypter(block, iv)
	// cfb := cipher.NewCBCEncrypter(block, iv)
	// cfb.CryptBlocks(ciphertext, text)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], text)

	// jklog.L().Infof("ciphertext: %v , len %d\n", ciphertext, len(ciphertext))

	return base64.StdEncoding.EncodeToString(ciphertext)
	// return hex.EncodeToString(ciphertext)
}

func JKAESDecrypt(key []byte, b64 string) string {
	text, err := base64.StdEncoding.DecodeString(b64)

	// text, err := hex.DecodeString(b64)
	if err != nil {
		jklog.L().Errorln("decode failed : ", err)
		return ""
	}

	// text := []byte(b64)
	jklog.L().Infoln("decode data len: ", len(text))

	block, err := aes.NewCipher(key)
	if err != nil {
		return ""
	}
	if len(text) < aes.BlockSize {
		return ""
	}

	// iv := text[:aes.BlockSize]
	// iv := []byte{
	// 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06,
	// 0x07, 0x08, 0x09, 0x0A, 0x0B,
	// 0x0C, 0x0D, 0x0E, 0x0F,
	// }
	iv := make([]byte, aes.BlockSize)
	// text = text[aes.BlockSize:]

	// jklog.L().Infoln("real data: ", text)

	if len(text)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	cfb := cipher.NewCFBDecrypter(block, iv)
	// cfb := cipher.NewCBCDecrypter(block, iv)
	// cfb.CryptBlocks(text, text)
	// jklog.L().Infof("first block %v\n", text)
	cfb.XORKeyStream(text, text)
	return string(text)
}
