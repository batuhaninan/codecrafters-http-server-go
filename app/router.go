package main

import (
	"os"
	"path/filepath"
)

var MAIN_ROUTE = NewRoute("/", GET, func(request RequestContext) Response {
	return Response{
		Status:  OK,
		Headers: []HttpHeader{},
		Body:    []byte("Hello World"),
	}
})

var ECHO_ROUTE = NewRoute("/echo/{}", GET, func(request RequestContext) Response {
	return Response{
		Status:  OK,
		Headers: []HttpHeader{},
		Body:    []byte(request.Params[0]),
	}
})

var USER_AGENT_ROUTE = NewRoute("/user-agent", GET, func(request RequestContext) Response {
	userAgent, ok := GetHeader(request.Headers, "User-Agent")

	if !ok {
		return Response{
			Status:  NOT_FOUND,
			Headers: []HttpHeader{},
		}
	}

	return Response{
		Status:  OK,
		Headers: []HttpHeader{},
		Body:    []byte(userAgent),
	}
})

var FILE_BY_ID_ROUTE = NewRoute("/files/{}", GET, func(request RequestContext) Response {
	filename := request.Params[0]

	if filename == "" {
		return Response{
			Status:  NOT_FOUND,
			Headers: []HttpHeader{},
		}
	}

	dir := "."

	os.ReadFile(filepath.Join(dir, filename))

	return Response{
		Status:  OK,
		Headers: []HttpHeader{NewHttpHeader("Content-Type", "application/octet-stream")},
		Body:    []byte(filename),
	}
})
