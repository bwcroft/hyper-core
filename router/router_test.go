package router

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type route struct {
	name       string
	path       string
	method     string
	reqPath    string
	exptVal    string
	exptCode   int
	skip       bool
	middleware []Middleware
}

type routeGroup struct {
	path       string
	routes     []route
	groups     []routeGroup
	middleware []Middleware
}

func makeRouteGroup() *routeGroup {
	return &routeGroup{
		routes: []route{
			{
				name:     "Root",
				path:     "/",
				reqPath:  "/",
				method:   http.MethodGet,
				exptVal:  "root",
				exptCode: http.StatusOK,
				middleware: []Middleware{
					NotFoundMiddleware,
				},
			},
			{
				name:     "Health",
				path:     "/health",
				reqPath:  "/health",
				method:   http.MethodGet,
				exptVal:  "healthy",
				exptCode: http.StatusOK,
			},
			{
				skip:     true,
				name:     "Invalid",
				path:     "/invalid",
				reqPath:  "/invalid",
				method:   http.MethodGet,
				exptCode: http.StatusNotFound,
			},
		},
		groups: []routeGroup{
			{
				path: "/api",
				routes: []route{
					{
						name:     "GetUsers",
						path:     "/users",
						reqPath:  "/api/users",
						method:   http.MethodGet,
						exptVal:  "got users",
						exptCode: http.StatusOK,
					},
					{
						name:     "GetUser",
						path:     "/users/{id}",
						reqPath:  "/api/users/1",
						method:   http.MethodGet,
						exptVal:  "got user",
						exptCode: http.StatusOK,
					},
					{
						name:     "PostUser",
						path:     "/users",
						reqPath:  "/api/users",
						method:   http.MethodPost,
						exptVal:  "created users",
						exptCode: http.StatusOK,
					},
					{
						name:     "PutUser",
						path:     "/users",
						reqPath:  "/api/users",
						method:   http.MethodPut,
						exptVal:  "put users",
						exptCode: http.StatusOK,
					},
					{
						name:     "PatchUser",
						path:     "/users",
						reqPath:  "/api/users",
						method:   http.MethodPatch,
						exptVal:  "patch users",
						exptCode: http.StatusOK,
					},
					{
						name:     "DeleteUser",
						path:     "/users",
						reqPath:  "/api/users",
						method:   http.MethodDelete,
						exptVal:  "delete users",
						exptCode: http.StatusOK,
					},
					{
						name:     "GetUserPosts",
						path:     "/users/{id}/posts",
						reqPath:  "/api/users/{id}/posts",
						method:   http.MethodGet,
						exptVal:  "user's posts",
						exptCode: http.StatusOK,
					},
				},
				groups: []routeGroup{
					{
						path:       "/protected",
						middleware: []Middleware{unauthorizedMiddleware},
						routes: []route{
							{
								name:     "GetProtectedVault",
								path:     "/vault",
								reqPath:  "/api/protected/vault",
								method:   http.MethodGet,
								exptVal:  "unauthorized",
								exptCode: http.StatusUnauthorized,
							},
						},
					},
				},
			},
		},
	}
}

func routeHandler(m string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, m)
	}
}

func unauthorizedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	})
}

func createTestRoutes(mux *Mux, rg *routeGroup) *Mux {
	for _, route := range rg.routes {
		if route.skip {
			continue
		}
		switch route.method {
		case http.MethodGet:
			mux.Get(route.path, routeHandler(route.exptVal), route.middleware...)
		case http.MethodPost:
			mux.Post(route.path, routeHandler(route.exptVal), route.middleware...)
		case http.MethodPut:
			mux.Put(route.path, routeHandler(route.exptVal), route.middleware...)
		case http.MethodPatch:
			mux.Patch(route.path, routeHandler(route.exptVal), route.middleware...)
		case http.MethodDelete:
			mux.Delete(route.path, routeHandler(route.exptVal), route.middleware...)
		default:
			panic(fmt.Sprintf("Invalid Route Method: route - %s method: %s", route.path, route.method))
		}
	}
	for _, nrg := range rg.groups {
		muxGroup := mux.Group(nrg.path, nrg.middleware...)
		createTestRoutes(muxGroup, &nrg)
	}
	return mux
}

func testRoute(t *testing.T, router *Mux, route *route) {
	t.Run(route.name, func(t *testing.T) {
		req, err := http.NewRequest(route.method, route.reqPath, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		if rr.Code != route.exptCode {
			t.Errorf("got reponse code %d, expected %d", rr.Code, route.exptCode)
		}
		body := strings.TrimSpace(rr.Body.String())
		if body != route.exptVal {
			t.Errorf("body contains %s, expected %s", body, route.exptVal)
		}
	})
}

func testRouteGroup(t *testing.T, router *Mux, group *routeGroup) {
	for _, route := range group.routes {
		testRoute(t, router, &route)
	}
	for _, group := range group.groups {
		testRouteGroup(t, router, &group)
	}
}

func TestRouter(t *testing.T) {
	r := New()
	rg := makeRouteGroup()
	createTestRoutes(r, rg)
	testRouteGroup(t, r, rg)
}
