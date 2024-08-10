package main

import (
	"fmt"
	"strings"
)

type Route struct {
	Path    string
	Method  HttpMethod
	Handler func(RequestMetadata) string
}

var ROUTES = make([]Route, 0)

var MAIN_ROUTE = NewRoute("/", GET, func(request RequestMetadata) string {
	return "Hello World"
})

var ECHO_ROUTE = NewRoute("/echo/{}", GET, func(request RequestMetadata) string {
	return request.Params[0]
})

var USER_AGENT_ROUTE = NewRoute("/user-agent", GET, func(request RequestMetadata) string {
	userAgent, ok := GetHeader(request.Headers, "User-Agent")

	if !ok {
		return "No User-Agent header"
	}

	return userAgent
})

func NewRoute(path string, method HttpMethod, handler func(RequestMetadata) string) Route {
	r := Route{
		Path:    path,
		Method:  method,
		Handler: handler,
	}

	ROUTES = append(ROUTES, r)

	return r
}

func GetRoute(request Request) (Route, RequestMetadata, error) {
	matchedMethodRoutes := Filter(ROUTES, func(r Route) bool {
		return r.Method == request.Line.HttpMethod
	})

	if len(matchedMethodRoutes) == 0 {
		return Route{}, RequestMetadata{}, fmt.Errorf("no route found for method %s", request.Line.HttpMethod)
	}

	routeSplit := DeleteEmptyStrings(strings.Split(request.Line.Target, "/"))

	for _, r := range matchedMethodRoutes {
		matchedRouteSplit := DeleteEmptyStrings(strings.Split(r.Path, "/"))

		if len(routeSplit) != len(matchedRouteSplit) {
			continue
		}

		fmt.Printf("Route Split: %+v\n", routeSplit)
		fmt.Printf("Matched Route Split: %+v\n", matchedRouteSplit)

		metadata := RequestMetadata{
			Headers: HttpHeadersToMap(request.Headers),
			Params:  make([]string, 0),
			Body:    request.Body,
		}

		found := false

		if request.Line.Target == "/" && r.Path == "/" {
			return r, metadata, nil
		}

		for i, s := range routeSplit {
			if s != matchedRouteSplit[i] {
				if matchedRouteSplit[i] == "{}" {
					metadata.Params = append(metadata.Params, s)
				} else {
					found = false
					break
				}
			} else {
				found = true
			}
		}

		if found {
			return r, metadata, nil
		}
	}

	return Route{}, RequestMetadata{}, fmt.Errorf("no route found for path %s", request.Line.Target)
}
