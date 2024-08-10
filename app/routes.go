package main

import (
	"fmt"
	"strings"
)

type Route struct {
	Path    string
	Method  HttpMethod
	Handler func(string, []string) string
}

var ROUTES = make([]Route, 0)

var MAIN_ROUTE = NewRoute("/", GET, func(body string, params []string) string {
	return "Hello World"
})

var ECHO_ROUTE = NewRoute("/echo/{}", GET, func(body string, params []string) string {
	return params[0]
})

func NewRoute(path string, method HttpMethod, handler func(string, []string) string) Route {
	r := Route{
		Path:    path,
		Method:  method,
		Handler: handler,
	}

	ROUTES = append(ROUTES, r)

	return r
}

func GetRoute(request Request) (Route, []string, error) {
	matchedMethodRoutes := Filter(ROUTES, func(r Route) bool {
		return r.Method == request.Line.HttpMethod
	})

	if len(matchedMethodRoutes) == 0 {
		return Route{}, nil, fmt.Errorf("no route found for method %s", request.Line.HttpMethod)
	}

	routeSplit := DeleteEmptyStrings(strings.Split(request.Line.Target, "/"))

	for _, r := range matchedMethodRoutes {
		matchedRouteSplit := DeleteEmptyStrings(strings.Split(r.Path, "/"))

		fmt.Printf("Route Split: %+v\n", routeSplit)
		fmt.Printf("Matched Route Split: %+v\n", matchedRouteSplit)

		if len(routeSplit) != len(matchedRouteSplit) {
			continue
		}

		params := make([]string, 0)

		for i, s := range routeSplit {
			if s != matchedRouteSplit[i] {
				if matchedRouteSplit[i] == "{}" {
					params = append(params, s)
				} else {
					break
				}
			}
		}

		return r, params, nil
	}

	return Route{}, nil, fmt.Errorf("no route found for path %s", request.Line.Target)
}
