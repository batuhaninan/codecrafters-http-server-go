package main

import (
	"bytes"
	"compress/gzip"
	"strings"
)

type Encoder func([]byte) ([]byte, error)

func Gzip(body []byte) ([]byte, error) {
	var b bytes.Buffer

	gz := gzip.NewWriter(&b)

	if _, err := gz.Write(body); err != nil {
		return nil, err
	}

	if err := gz.Close(); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

var VALID_ENCODINGS = map[string]Encoder{
	"gzip": Gzip,
}

func GetEncoder(key string) (Encoder, bool) {
	for k, v := range VALID_ENCODINGS {
		if strings.EqualFold(k, key) {
			return v, true
		}
	}

	return nil, false
}
