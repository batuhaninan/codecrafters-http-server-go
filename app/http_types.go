package main

import "fmt"

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
		if string(m) == method {
			return m, nil
		}
	}

	return "", fmt.Errorf("unknown http method: %s", method)
}

func GetHttpVersion(version string) (HttpVersion, error) {
	for _, m := range HTTP_VERSIONS {
		if string(m) == version {
			return m, nil
		}
	}

	return "", fmt.Errorf("unknown http version: %s", version)
}
