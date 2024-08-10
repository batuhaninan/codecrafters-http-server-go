package main

type HttpStatus struct {
	StatusCode int
	Reason     string
}

var OK = HttpStatus{
	StatusCode: 200,
	Reason:     "OK",
}

var NOT_FOUND = HttpStatus{
	StatusCode: 404,
	Reason:     "Not Found",
}

var BAD_REQUEST = HttpStatus{
	StatusCode: 400,
	Reason:     "Bad Request",
}

var INTERNAL_SERVER_ERROR = HttpStatus{
	StatusCode: 500,
	Reason:     "Internal Server Error",
}

var CREATED = HttpStatus{
	StatusCode: 201,
	Reason:     "Created",
}
