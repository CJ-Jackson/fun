package commonUtil

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
	"log"
)

func DecodeValue(value string) []byte {
	raw, err := base64.StdEncoding.DecodeString(value)
	errCheck(err)

	r, err := gzip.NewReader(bytes.NewReader(raw))
	errCheck(err)

	buf := &bytes.Buffer{}
	_, err = io.Copy(buf, r)
	errCheck(err)

	return buf.Bytes()
}

func DecodeValueStr(value string) string { return string(DecodeValue(value)) }

func errCheck(err error) {
	if err != nil {
		log.Panic(err)
	}
}
