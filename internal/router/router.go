package router

import (
	"errors"
	"reverse-proxy/internal/config"
)

type Router struct {
	routes map[string]string
}

func New(routes []config.Route) (*Router, error) {
	if len(routes) < 1 {
		return nil, errors.New("can't create empty Router")
	}

	router := &Router{
		routes: make(map[string]string),
	}

	for _, r := range routes {
		router.routes[r.Prefix] = r.Upstream
	}

	return router, nil
}

func (r *Router) Match(path string) (upstream string, ok bool) {
	upstream, ok = r.routes[path]
	return upstream, ok
}
