package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bwcroft/hyper-core/internal/config"
	"github.com/bwcroft/hyper-core/internal/database/models"
	"github.com/bwcroft/hyper-core/utils"
	"github.com/jackc/pgx/v5"
)

func name() {}

func StartServer(db *pgx.Conn) (err error) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		var user models.User
		query := "SELECT id, first_name, middle_name, last_name, email, phone, created_at, updated_at FROM USERS WHERE id = $1"
		err := db.QueryRow(context.Background(), query, 1).Scan(
			&user.ID,
			&user.FirstName,
			&user.MiddleName,
			&user.LastName,
			&user.Email,
			&user.Phone,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			utils.LogError(&err)
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, fmt.Sprintf("Unable to encode JSON: %v", err), http.StatusInternalServerError)
			return
		}
	})

	portNum := utils.GetEnvUint16(config.ServerPort, 8080)
	port := fmt.Sprintf(":%d", portNum)

	fmt.Printf("Server started on port %s\n", port)
	err = http.ListenAndServe(port, mux)
	return
}
