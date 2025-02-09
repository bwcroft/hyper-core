package handlers

import (
	"fmt"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Get User")
  // if r.URL.Path != "/" {
  //   http.NotFound(w, r)
  //   return
  // }

  // var user models.User
  // query := "SELECT id, first_name, middle_name, last_name, email, phone, created_at, updated_at FROM USERS WHERE id = $1"
  // if err := c.DB.Get(&user, query, 1); err != nil {
  //   utils.LogError(err)
  // }

  // w.Header().Set("Content-Type", "application/json")
  // if err := json.NewEncoder(w).Encode(user); err != nil {
  //   http.Error(w, fmt.Sprintf("Unable to encode JSON: %v", err), http.StatusInternalServerError)
  //   return
  // }
}
