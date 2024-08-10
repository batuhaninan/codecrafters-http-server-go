package main

import "strings"

type HttpHeader struct {
	Key   string
	Value string
}

func GetHeader(headers []HttpHeader, key string) (string, bool) {
	for _, h := range headers {
		if strings.EqualFold(h.Key, key) {
			return h.Value, true
		}
	}
	return "", false
}

func HasHeader(headers []HttpHeader, key string) bool {
	for _, h := range headers {
		if strings.EqualFold(h.Key, key) {
			return true
		}
	}
	return false
}

func HttpHeadersToMap(headers []HttpHeader) map[string]string {
	m := make(map[string]string)

	for _, h := range headers {
		m[h.Key] = h.Value
	}

	return m
}

func NewHttpHeader(key string, value string) HttpHeader {
	return HttpHeader{
		Key:   key,
		Value: value,
	}
}
