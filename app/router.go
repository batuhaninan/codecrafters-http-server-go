package main

import (
	"os"
	"path/filepath"
	"strings"
)

var MAIN_ROUTE = NewRoute("/", GET, func(ctx RequestContext) Response {
	return Response{
		Status:  OK,
		Headers: []HttpHeader{},
		Body:    []byte("Hello World"),
	}
})

var ECHO_ROUTE = NewRoute("/echo/{}", GET, func(ctx RequestContext) Response {

	encoding, ok := GetHeader(ctx.Headers, "Accept-Encoding")

	headers := []HttpHeader{}

	encodings := DeleteEmptyStrings(strings.Split(encoding, ", "))

	if ok {
		for _, encoding := range encodings {
			if Contains(VALID_ENCODINGS, encoding) {
				headers = append(headers, NewHttpHeader("Content-Encoding", encoding))
			}
		}
	}

	return Response{
		Status:  OK,
		Headers: headers,
		Body:    []byte(ctx.Params[0]),
	}
})

var USER_AGENT_ROUTE = NewRoute("/user-agent", GET, func(ctx RequestContext) Response {
	userAgent, ok := GetHeader(ctx.Headers, "User-Agent")

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

var FILE_BY_ID_ROUTE = NewRoute("/files/{}", GET, func(ctx RequestContext) Response {
	filename := ctx.Params[0]

	if filename == "" {
		return Response{
			Status:  NOT_FOUND,
			Headers: []HttpHeader{},
		}
	}

	fileContent, err := os.ReadFile(filepath.Join(ctx.ServerOpts.ServeDirectory, filename))

	if err != nil {
		return Response{
			Status:  NOT_FOUND,
			Headers: []HttpHeader{},
		}
	}

	return Response{
		Status:  OK,
		Headers: []HttpHeader{NewHttpHeader("Content-Type", "application/octet-stream")},
		Body:    fileContent,
	}
})

var FILE_CREATE_ROUTE = NewRoute("/files/{}", POST, func(ctx RequestContext) Response {
	filename := ctx.Params[0]

	if filename == "" {
		return Response{
			Status:  NOT_FOUND,
			Headers: []HttpHeader{},
		}
	}

	file, err := os.Create(filepath.Join(ctx.ServerOpts.ServeDirectory, filename))

	if err != nil {
		return Response{
			Status:  INTERNAL_SERVER_ERROR,
			Headers: []HttpHeader{},
		}
	}
	defer file.Close()

	file.Write(ctx.Body)

	return Response{
		Status: CREATED,
	}
})
