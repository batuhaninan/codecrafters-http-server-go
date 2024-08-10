package main

import (
	"fmt"
	"net"
	"strings"
)

type Response struct {
	Status  HttpStatus
	Headers []HttpHeader
	Body    []byte
}

func (s *Server) sendResponse(conn net.Conn, response Response) {

	headers := response.Headers

	body := response.Body

	if !HasHeader(headers, "Content-Type") {
		headers = append(headers, NewHttpHeader("Content-Type", "text/plain"))
	}

	if encoding, ok := GetHeader(headers, "Content-Encoding"); ok {

		encoder, ok := GetEncoder(encoding)

		if ok {
			newBody, err := encoder(body)

			if err == nil {
				body = newBody
			}
		}
	}

	bodyFinal := body

	headers = append(headers, NewHttpHeader("Content-Length", fmt.Sprintf("%d", len(bodyFinal))))

	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("%s %d %s\r\n", s.HttpOpts.Version, response.Status.StatusCode, response.Status.Reason))

	for _, h := range headers {
		sb.WriteString(fmt.Sprintf("%s: %s\r\n", h.Key, h.Value))
	}

	sb.WriteString("\r\n")

	sbBytes := []byte(sb.String())

	sbBytes = append(sbBytes, bodyFinal...)

	conn.Write(sbBytes)
}
