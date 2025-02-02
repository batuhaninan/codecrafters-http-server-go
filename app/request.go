package main

import (
	"errors"
	"strings"
)

type Request struct {
	Line    RequestLine
	Headers []HttpHeader
	Body    []byte
}

type RequestContext struct {
	Headers    []HttpHeader
	Params     []string
	Body       []byte
	ServerOpts ServerOpts
}

func parseRequestLine(request string) (RequestLine, error) {
	line := strings.Split(request, " ")

	if len(line) < 3 {
		return RequestLine{}, errors.New("invalid request")
	}

	method, err := GetHttpMethod(line[0])

	if err != nil {
		return RequestLine{}, err
	}

	version, err := GetHttpVersion(line[2])

	if err != nil {
		return RequestLine{}, err
	}

	return RequestLine{
		HttpMethod: method,
		Target:     line[1],
		Version:    version,
	}, nil
}

func parseHeaders(request string) ([]HttpHeader, []byte, error) {
	firstCRLF := strings.Index(request, "\r\n")

	headerStart := request[firstCRLF+2:]

	rest := strings.Split(headerStart, "\r\n\r\n")

	body := make([]byte, 0)

	if len(rest) > 1 {
		body = []byte(rest[1])
	}

	_headers := DeleteEmptyStrings(strings.Split(rest[0], "\r\n"))

	headers := make([]HttpHeader, 0)

	for _, h := range _headers {
		if len(h) == 0 {
			continue
		}

		headerParts := DeleteEmptyStrings(strings.Split(h, ": "))

		if len(headerParts) != 2 {
			return []HttpHeader{}, []byte{}, errors.New("invalid header")
		}

		headers = append(headers, HttpHeader{
			Key:   headerParts[0],
			Value: headerParts[1],
		})
	}

	return headers, body, nil
}

func parseRequest(request string) (Request, error) {

	parts := strings.Split(request, "\r\n")

	if len(parts) < 1 {
		return Request{}, errors.New("invalid request")
	}

	reqLine, err := parseRequestLine(parts[0])

	if err != nil {
		return Request{}, err
	}

	headers, body, err := parseHeaders(request)

	if err != nil {
		return Request{}, err
	}

	req := Request{
		Line:    reqLine,
		Headers: headers,
		Body:    body,
	}

	return req, nil
}
