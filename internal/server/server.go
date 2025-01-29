package server

import (
	"fmt"
	"net/http"

	"github.com/bwcroft/hyper-core/internal/config"
	"github.com/bwcroft/hyper-core/utils"
)

func name() {}

func StartServer() (err error) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Write([]byte("HyperCore Server"))
	})

	portNum := utils.GetEnvUint16(config.ServerPort, 8080)
	port := fmt.Sprintf(":%d", portNum)

	fmt.Printf("Server started on port %s\n", port)
	err = http.ListenAndServe(port, mux)
	return
}
