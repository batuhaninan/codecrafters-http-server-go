package main

import (
	"fmt"
	"net"
)

func (s *Server) sendResponse(conn net.Conn, status HttpStatus) {
	conn.Write([]byte(fmt.Sprintf("%s %d %s\r\n\r\n", s.HttpOpts.Version, status.StatusCode, status.Reason)))
}
