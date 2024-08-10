package main

import (
	"fmt"
	"strings"
)

type HttpVersion string
type HttpMethod string

type RequestLine struct {
	HttpMethod HttpMethod
	Target     string
	Version    HttpVersion
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
