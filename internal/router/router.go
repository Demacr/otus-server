package router

import (
	"net/http"

	storages "../storages"
)

// Router - struct which contain rootHandler only
type Router struct {
	rootHandler rootHandler
}

// New - generates new Router object
func New(store storages.Store) *Router {
	return &Router{
		rootHandler: newRootHandler(store),
	}
}

// RootHandler handler of the root path
func (r *Router) RootHandler() http.Handler {
	return r.rootHandler
}
