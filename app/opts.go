package main

type ServerOpts struct {
	Port           int
	ServeDirectory string
}

type HttpOpts struct {
	Version HttpVersion
}
