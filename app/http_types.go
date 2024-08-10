package main

import (
	"fmt"
	"strings"
)

type HttpVersion string
type HttpMethod string

type HttpStatus struct {
	StatusCode int
	Reason     string
}

type RequestLine struct {
	HttpMethod HttpMethod
	Target     string
	Version    HttpVersion
}

type HttpHeader struct {
	Key   string
	Value string
}

var OK = HttpStatus{
	StatusCode: 200,
	Reason:     "OK",
}

var NOT_FOUND = HttpStatus{
	StatusCode: 404,
	Reason:     "Not Found",
}

var BAD_REQUEST = HttpStatus{
	StatusCode: 400,
	Reason:     "Bad Request",
}

var INTERNAL_SERVER_ERROR = HttpStatus{
	StatusCode: 500,
	Reason:     "Internal Server Error",
}

var CREATED = HttpStatus{
	StatusCode: 201,
	Reason:     "Created",
}

var (
	HTTP_1_1 HttpVersion = "HTTP/1.1"
)

var HTTP_VERSIONS = []HttpVersion{
	HTTP_1_1,
}

var (
	GET     HttpMethod = "GET"
	POST    HttpMethod = "POST"
	PUT     HttpMethod = "PUT"
	PATCH   HttpMethod = "PATCH"
	DELETE  HttpMethod = "DELETE"
	OPTIONS HttpMethod = "OPTIONS"
	HEAD    HttpMethod = "HEAD"
)

var HTTP_METHODS = []HttpMethod{
	GET,
	POST,
	PUT,
	PATCH,
	DELETE,
	OPTIONS,
	HEAD,
}

func HasHeader(headers []HttpHeader, key string) bool {
	for _, h := range headers {
		if strings.EqualFold(h.Key, key) {
			return true
		}
	}
	return false
}

func GetHeader(headers map[string]string, key string) (string, bool) {
	for k, v := range headers {
		if strings.EqualFold(k, key) {
			return v, true
		}
	}

	return "", false
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

func GetHttpMethod(method string) (HttpMethod, error) {
	for _, m := range HTTP_METHODS {
		if string(m) == strings.ToUpper(method) {
			return m, nil
		}
	}

	return "", fmt.Errorf("unknown http method: %s", method)
}

func GetHttpVersion(version string) (HttpVersion, error) {
	for _, m := range HTTP_VERSIONS {
		if string(m) == strings.ToUpper(version) {
			return m, nil
		}
	}

	return "", fmt.Errorf("unknown http version: %s", version)
}
