package routes

import (
	"fmt"
	"net/http"

	"github.com/bwcroft/hyper-core/internal/handlers"
	"github.com/bwcroft/hyper-core/router"
	"github.com/jmoiron/sqlx"
)

func InitRoutes(db *sqlx.DB) (routes *router.Mux) {
	routes = router.New()
	routes.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hyper Core")
	})
	routes.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Health Check")
	})

	v1 := routes.Group("/v1")
	users := v1.Group("/users")
	users.Get("/", handlers.GetUser)
	users.Post("/", handlers.GetUser)
	users.Put("/", handlers.GetUser)
	users.Patch("/", handlers.GetUser)
	users.Delete("/", handlers.GetUser)
	return
}
