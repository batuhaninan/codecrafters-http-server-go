package main

import (
	"fmt"
	"net"
	"strings"
)

type Response struct {
	Status  HttpStatus
	Headers []HttpHeader
	Body    string
}

func (s *Server) sendResponse(conn net.Conn, response Response) {

	headers := response.Headers

	headers = append(headers, NewHttpHeader("Content-Length", fmt.Sprintf("%d", len(response.Body))))
	headers = append(headers, NewHttpHeader("Content-Type", "text/plain"))

	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("%s %d %s\r\n", s.HttpOpts.Version, response.Status.StatusCode, response.Status.Reason))

	for _, h := range headers {
		sb.WriteString(fmt.Sprintf("%s: %s\r\n", h.Key, h.Value))
	}

	sb.WriteString("\r\n")
	sb.WriteString(response.Body)

	fmt.Printf("Sending response: %q\n", sb.String())

	conn.Write([]byte(
		sb.String(),
	),
	)
}
