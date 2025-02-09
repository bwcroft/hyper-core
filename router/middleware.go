package router

import (
	"fmt"
	"net/http"
	// "strings"
	"time"

	"github.com/bwcroft/hyper-core/utils"
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

func StackMiddlerware(m ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(m) - 1; i >= 0; i-- {
			next = m[i](next)
		}
		return next
	}
}

func RequestLog(next http.Handler) http.Handler {
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

func ValidatePath(routes map[string]HttpMethod) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqRoute := r.URL.Path
			isValid := false
			for route := range routes {
        fmt.Printf("ReqRoute: %s, Route: %s\n", reqRoute, route)
				if reqRoute == route{
					isValid = true
					break
				}
        //TODO: If not exact match, check that we have a match with a unique url param
			}
			if !isValid {
				http.NotFound(w, r)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
