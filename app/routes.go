package main

import (
	"fmt"
	"strings"
)

type Route struct {
	Path    string
	Method  HttpMethod
	Handler func(RequestContext) Response
}

var ROUTES = make([]Route, 0)

func NewRoute(path string, method HttpMethod, handler func(RequestContext) Response) Route {
	r := Route{
		Path:    path,
		Method:  method,
		Handler: handler,
	}

	ROUTES = append(ROUTES, r)

	return r
}

func (s *Server) GetRoute(request Request) (Route, RequestContext, error) {
	matchedMethodRoutes := Filter(ROUTES, func(r Route) bool {
		return r.Method == request.Line.HttpMethod
	})

	if len(matchedMethodRoutes) == 0 {
		return Route{}, RequestContext{}, fmt.Errorf("no route found for method %s", request.Line.HttpMethod)
	}

	routeSplit := DeleteEmptyStrings(strings.Split(request.Line.Target, "/"))

	for _, r := range matchedMethodRoutes {
		matchedRouteSplit := DeleteEmptyStrings(strings.Split(r.Path, "/"))

		if len(routeSplit) != len(matchedRouteSplit) {
			continue
		}

		metadata := RequestContext{
			Headers:    HttpHeadersToMap(request.Headers),
			Params:     make([]string, 0),
			Body:       request.Body,
			ServerOpts: s.Opts,
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

	return Route{}, RequestContext{}, fmt.Errorf("no route found for path %s", request.Line.Target)
}
