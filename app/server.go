package main

import (
	"fmt"
	"log/slog"
	"net"
	"os"
)

type ServerOpts struct {
	Port int
}

type HttpOpts struct {
	Version string
}

type Server struct {
	Opts     ServerOpts
	HttpOpts HttpOpts
}

type HttpStatus struct {
	StatusCode int
	Reason     string
}

var OK = HttpStatus{
	StatusCode: 200,
	Reason:     "OK",
}

func main() {
	fmt.Println("Server started")

	server := Server{
		Opts: ServerOpts{
			4221,
		},
		HttpOpts: HttpOpts{
			Version: "1.1",
		},
	}

	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", server.Opts.Port))
	if err != nil {
		slog.Error("Failed to bind to port:", "port", server.Opts.Port)
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		slog.Error("Error accepting connection:", "error", err.Error())
		os.Exit(1)
	}

	conn.Write([]byte(fmt.Sprintf("HTTP/%s %d %s\r\n\r\n", server.HttpOpts.Version, OK.StatusCode, OK.Reason)))
}
