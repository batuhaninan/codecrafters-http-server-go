package main

import (
	"fmt"
	"log/slog"
	"net"
	"os"
)

type Server struct {
	Opts     ServerOpts
	HttpOpts HttpOpts
}

func (s *Server) readLoop(conn net.Conn) {

	for {
		bytes := make([]byte, 1024)

		nbytes, err := conn.Read(bytes)

		if err != nil {
			continue
		}

		got := bytes[:nbytes]

		request, err := parseRequest(string(got))

		if err != nil {
			fmt.Printf("Error parsing request: %s\n", err.Error())
			s.sendResponse(conn, Response{Status: BAD_REQUEST})
			continue
		}

		fmt.Printf("Request: %+v\n", request)

		if route, metadata, err := GetRoute(request); err == nil {
			response := Response{
				Status:  OK,
				Headers: []HttpHeader{},
				Body:    route.Handler(metadata),
			}
			s.sendResponse(conn, response)
		} else {
			s.sendResponse(conn, Response{Status: NOT_FOUND})
		}
	}

}

func main() {
	fmt.Println("Server started")

	server := Server{
		Opts: ServerOpts{
			4221,
		},
		HttpOpts: HttpOpts{
			Version: HTTP_1_1,
		},
	}

	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", server.Opts.Port))
	if err != nil {
		slog.Error("Failed to bind to port:", "port", server.Opts.Port)
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			slog.Error("Error accepting connection:", "error", err.Error())
			os.Exit(1)
		}

		go server.readLoop(conn)
	}

}
