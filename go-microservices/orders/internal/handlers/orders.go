package handlers

import (
"encoding/json"
"net/http"

"orders/internal/service"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(service.GetAll())
}
