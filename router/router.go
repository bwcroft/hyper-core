package router

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/bwcroft/hyper-core/utils"
)

type HttpMethod string

const (
	HttpGet    HttpMethod = "GET"
	HttpPost   HttpMethod = "POST"
	HttpPut    HttpMethod = "PUT"
	HttpPatch  HttpMethod = "PATCH"
	HttpDelete HttpMethod = "DELETE"
)

type Mux struct {
	*http.ServeMux

	/** Route path is the key and request type is the value */
	routes map[string]HttpMethod

	/** If group path if the mux is apart of one */
	groupPath string

	/** The parent mux of the children */
	parent *Mux
}

func New() *Mux {
	return &Mux{
		ServeMux: http.NewServeMux(),
		routes:   make(map[string]HttpMethod),
	}
}

func (mux *Mux) createHandler(method HttpMethod, path string, fn http.HandlerFunc, m ...Middleware) {
	route := fmt.Sprintf("%s%s", mux.groupPath, path)
	if mux.parent != nil && mux.parent.routes != nil {
		(mux.parent.routes)[route] = method
	} else if mux.routes != nil {
		(mux.routes)[route] = method
	}
	stack := StackMiddlerware(m...)
	mux.Handle(fmt.Sprintf("%s %s", method, path), stack(fn))
}

func (mux *Mux) RoutesList() (rs []string) {
	if mux.routes != nil {
    //TODO: Fix missing routes
    for k, v := range mux.routes {
      rs = append(rs, fmt.Sprintf("%s %s", k, v))
    }
	}
	sort.Strings(rs)
	return
}

func (mux *Mux) Get(path string, fn http.HandlerFunc, m ...Middleware) {
	mux.createHandler(HttpGet, path, fn, m...)
}

func (mux *Mux) Post(path string, fn http.HandlerFunc, m ...Middleware) {
	mux.createHandler(HttpPost, path, fn, m...)
}

func (mux *Mux) Put(path string, fn http.HandlerFunc, m ...Middleware) {
	mux.createHandler(HttpPut, path, fn, m...)
}

func (mux *Mux) Patch(path string, fn http.HandlerFunc, m ...Middleware) {
	mux.createHandler(HttpPatch, path, fn, m...)
}

func (mux *Mux) Delete(path string, fn http.HandlerFunc, m ...Middleware) {
	mux.createHandler(HttpDelete, path, fn, m...)
}

func (mux *Mux) Group(path string, m ...Middleware) (cmux *Mux) {
	cmux = &Mux{
		ServeMux:  http.NewServeMux(),
		groupPath: fmt.Sprintf("%s%s", mux.groupPath, path),
	}
	if mux.parent != nil {
		cmux.parent = mux.parent
	} else {
		cmux.parent = mux
	}

	stack := StackMiddlerware(m...)
	httpMethods := [5]HttpMethod{HttpGet, HttpPost, HttpPut, HttpPatch, HttpDelete}
	for _, m := range httpMethods {
		mux.Handle(fmt.Sprintf("%s %s/", m, path), http.StripPrefix(path, stack(cmux)))
	}
	return
}

func (mux *Mux) Listen(port uint16) (err error) {
  stack := StackMiddlerware(
    RequestLog,
    ValidatePath(mux.routes),
  )
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: stack(mux),
	}

	utils.InfoBox(append([]string{
    "HyperCore: 1.0.0",
	  fmt.Sprintf("Server Port %d", port),
    "Routes:",
	}, mux.RoutesList()...))

	err = server.ListenAndServe()
  return
}
