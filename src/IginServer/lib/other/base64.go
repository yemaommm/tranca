package other

import (
	"encoding/base64"
	// "fmt"
)

var Base = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func Base64Encode(src []byte, base64table string) []byte {
	var coder *base64.Encoding
	if base64table == "" {
		coder = base64.NewEncoding(Base)
	} else {
		coder = base64.NewEncoding(base64table)
	}
	return []byte(coder.EncodeToString(src))
}

func Base64Decode(src []byte, base64table string) ([]byte, error) {
	var coder *base64.Encoding
	if base64table == "" {
		coder = base64.NewEncoding(Base)
	} else {
		coder = base64.NewEncoding(base64table)
	}
	return coder.DecodeString(string(src))
}
