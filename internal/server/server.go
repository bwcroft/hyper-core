package server

import (
	"fmt"
	"net/http"

	"github.com/bwcroft/hyper-core/internal/config"
	"github.com/bwcroft/hyper-core/utils"
)

func StartServer() (err error) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("HyperCore Server"))
	})
	port := utils.GetEnvUint16(config.ServerPort, 8080)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
  return
}
