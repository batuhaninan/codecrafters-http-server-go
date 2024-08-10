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

	if !HasHeader(headers, "Content-Type") {
		headers = append(headers, NewHttpHeader("Content-Type", "text/plain"))
	}
	headers = append(headers, NewHttpHeader("Content-Length", fmt.Sprintf("%d", len(response.Body))))

	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("%s %d %s\r\n", s.HttpOpts.Version, response.Status.StatusCode, response.Status.Reason))

	for _, h := range headers {
		sb.WriteString(fmt.Sprintf("%s: %s\r\n", h.Key, h.Value))
	}

	sb.WriteString("\r\n")
	// sb.WriteString(response.Body)

	sbBytes := []byte(sb.String())

	sbBytes = append(sbBytes, response.Body...)

	conn.Write(sbBytes)
}
