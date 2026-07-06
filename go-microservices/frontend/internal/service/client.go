package service

import (
"encoding/json"
"net/http"

"frontend/internal/models"
)

func fetch(url string) models.BackendResponse {
resp, err := http.Get(url)
if err != nil {
return models.BackendResponse{Service: "error", Data: []string{err.Error()}}
}
defer resp.Body.Close()

var result models.BackendResponse
json.NewDecoder(resp.Body).Decode(&result)

return result
}

func LoadPageData() models.PageData {
return models.PageData{
Users:  fetch("http://users-service:8080/users"),
Orders: fetch("http://orders-service:8080/orders"),
Cart:   fetch("http://cart-service:8080/cart"),
}
}
