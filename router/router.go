package router

import (
	"fmt"
	"net/http"

	"github.com/bwcroft/hyper-core/utils"
)

type Mux struct {
	*http.ServeMux
	parent           *Mux
	prefix           string
	prefixMethods    map[string]bool
	prefixMiddleware *[]Middleware
}

func New() *Mux {
	return &Mux{
		ServeMux: http.NewServeMux(),
	}
}

func (mux *Mux) createPrefix(method string) {
	if mux.parent == nil || mux.prefix == "" || mux.prefixMethods == nil || mux.prefixMethods[method] {
		return
	}
	stack := StackMiddleware(*mux.prefixMiddleware...)
	fmt.Printf("Prefix: %s %s/\n", method, mux.prefix)
	mux.parent.Handle(fmt.Sprintf("%s %s/", method, mux.prefix), http.StripPrefix(mux.prefix, stack(mux)))
	mux.prefixMethods[method] = true
	return
}

func (mux *Mux) createHandler(method string, path *string, fn *http.HandlerFunc, m *[]Middleware) {
	mux.createPrefix(method)
	stack := StackMiddleware(*m...)
	mux.Handle(fmt.Sprintf("%s %s", method, *path), stack(fn))
}

func (mux *Mux) Get(path string, fn http.HandlerFunc, m ...Middleware) {
	mux.createHandler(http.MethodGet, &path, &fn, &m)
}

func (mux *Mux) Post(path string, fn http.HandlerFunc, m ...Middleware) {
	mux.createHandler(http.MethodPost, &path, &fn, &m)
}

func (mux *Mux) Put(path string, fn http.HandlerFunc, m ...Middleware) {
	mux.createHandler(http.MethodPut, &path, &fn, &m)
}

func (mux *Mux) Patch(path string, fn http.HandlerFunc, m ...Middleware) {
	mux.createHandler(http.MethodPatch, &path, &fn, &m)
}

func (mux *Mux) Delete(path string, fn http.HandlerFunc, m ...Middleware) {
	mux.createHandler(http.MethodDelete, &path, &fn, &m)
}

func (mux *Mux) Connect(path string, fn http.HandlerFunc, m ...Middleware) {
	mux.createHandler(http.MethodConnect, &path, &fn, &m)
}

func (mux *Mux) Options(path string, fn http.HandlerFunc, m ...Middleware) {
	mux.createHandler(http.MethodOptions, &path, &fn, &m)
}

func (mux *Mux) Group(prefix string, m ...Middleware) *Mux {
	return &Mux{
		ServeMux:         http.NewServeMux(),
		prefix:           prefix,
		prefixMethods:    make(map[string]bool),
		prefixMiddleware: &m,
		parent:           mux,
	}
}

func (mux *Mux) Listen(port uint16) (err error) {
	stack := StackMiddleware(LoggerMiddleware)
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: stack(mux),
	}
	utils.InfoBox([]string{
		fmt.Sprintf("HyperCore - running on port %d", port),
	})
	err = server.ListenAndServe()
	return
}
