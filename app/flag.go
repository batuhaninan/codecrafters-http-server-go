package main

import (
	"flag"
)

type Args struct {
	Directory string
}

func InitArgs() Args {

	var directoryFlag string
	flag.StringVar(&directoryFlag, "directory", ".", "Directory to serve")

	flag.Parse()

	return Args{
		Directory: directoryFlag,
	}
}
