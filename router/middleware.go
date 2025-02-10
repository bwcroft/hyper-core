package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bwcroft/hypercore/utils"
)

type Middleware func(http.Handler) http.Handler

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

// StackMiddleware chains multiple middleware functions 
func StackMiddleware(m *[]Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(*m) - 1; i >= 0; i-- {
			next = (*m)[i](next)
		}
		return next
	}
}

// LoggerMiddleware is responsible for logging details of each incoming request, 
// including the method, URL, and timestamp. It helps track and monitor requests 
// made to the server for debugging and performance analysis.
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		start := time.Now()

		w := &wrappedWriter{rw, http.StatusOK}
		next.ServeHTTP(w, r)

		if r.URL.Path != "/health" {
			msg := fmt.Sprintf("%d %s %s %v", w.statusCode, r.Method, r.URL.Path, time.Since(start))
			utils.LogInfo(msg)
		}
	})
}

// NotFoundMiddleware prevents unknown paths from falling back to the base route 
// in the standard HTTP router. It checks if the requested path is valid and 
// returns a 404 Not Found response for unrecognized paths.
// Use only with base HTTP methods and the root path ("/"), such as GET /, POST /, etc.
func NotFoundMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
		  http.Error(w, "", http.StatusNotFound)
      return
    } 
    next.ServeHTTP(w, r)
	})
}
