package routes

import (
	"fmt"
	"net/http"

	"github.com/bwcroft/hyper-core/internal/handlers"
	"github.com/bwcroft/hyper-core/router"
	"github.com/jmoiron/sqlx"
)

func InitRoutes(db *sqlx.DB) (routes *router.Mux) {
	routes = router.New(
    router.LoggerMiddleware,
  )

	routes.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hyper Core")
	})
	routes.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Health Check")
	})

	v1 := routes.Group("/v1")
	v1.Get("/users", handlers.GetUser)
	v1.Post("/users", handlers.GetUser)
	v1.Put("/users", handlers.GetUser)
	v1.Patch("/users", handlers.GetUser)
	v1.Delete("/users", handlers.GetUser)
	return
}
