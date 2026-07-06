# micro-services-kubernetes-go


$root = "go-microservices"

$dirs = @(
  "$root/frontend/cmd/frontend",
  "$root/frontend/internal/handlers",
  "$root/frontend/internal/service",
  "$root/frontend/internal/models",
  "$root/frontend/internal/templates",

  "$root/users/cmd/users",
  "$root/users/internal/handlers",
  "$root/users/internal/service",
  "$root/users/internal/repository",
  "$root/users/internal/models",

  "$root/orders/cmd/orders",
  "$root/orders/internal/handlers",
  "$root/orders/internal/service",
  "$root/orders/internal/repository",
  "$root/orders/internal/models",

  "$root/cart/cmd/cart",
  "$root/cart/internal/handlers",
  "$root/cart/internal/service",
  "$root/cart/internal/repository",
  "$root/cart/internal/models",

  "$root/k8s/frontend",
  "$root/k8s/users",
  "$root/k8s/orders",
  "$root/k8s/cart",
  "$root/k8s/gateway",
  "$root/pkg"
)

foreach ($dir in $dirs) {
  New-Item -ItemType Directory -Force -Path $dir | Out-Null
}

@"
go 1.22

use (
    ./frontend
    ./users
    ./orders
    ./cart
)
"@ | Set-Content "$root/go.work"

$services = @("frontend", "users", "orders", "cart")

foreach ($svc in $services) {
@"
module $svc

go 1.22
"@ | Set-Content "$root/$svc/go.mod"

@"
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o app ./cmd/$svc

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]
"@ | Set-Content "$root/$svc/Dockerfile"
}

@"
package main

import (
	"log"
	"net/http"

	"frontend/internal/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Home)

	log.Println("frontend running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
"@ | Set-Content "$root/frontend/cmd/frontend/main.go"

@"
package handlers

import (
	"html/template"
	"net/http"

	"frontend/internal/service"
)

func Home(w http.ResponseWriter, r *http.Request) {
	data := service.LoadPageData()

	tmpl := template.Must(template.ParseFiles("internal/templates/index.html"))
	tmpl.Execute(w, data)
}
"@ | Set-Content "$root/frontend/internal/handlers/home.go"

@"
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
"@ | Set-Content "$root/frontend/internal/service/client.go"

@"
package models

type BackendResponse struct {
	Service string   `json:"service"`
	Data    []string `json:"data"`
}

type PageData struct {
	Users  BackendResponse
	Orders BackendResponse
	Cart   BackendResponse
}
"@ | Set-Content "$root/frontend/internal/models/response.go"

@"
<!DOCTYPE html>
<html>
<head>
	<title>Go Microservices</title>
</head>
<body>
	<h1>Go Microservices Frontend</h1>

	<h2>Users</h2>
	<ul>{{range .Users.Data}}<li>{{.}}</li>{{end}}</ul>

	<h2>Orders</h2>
	<ul>{{range .Orders.Data}}<li>{{.}}</li>{{end}}</ul>

	<h2>Cart</h2>
	<ul>{{range .Cart.Data}}<li>{{.}}</li>{{end}}</ul>
</body>
</html>
"@ | Set-Content "$root/frontend/internal/templates/index.html"

function Create-Backend {
  param (
    [string]$name,
    [string]$path,
    [string]$items
  )

@"
package main

import (
	"log"
	"net/http"

	"$name/internal/handlers"
)

func main() {
	http.HandleFunc("/$path", handlers.GetAll)

	log.Println("$name service running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
"@ | Set-Content "$root/$name/cmd/$name/main.go"

@"
package handlers

import (
	"encoding/json"
	"net/http"

	"$name/internal/service"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(service.GetAll())
}
"@ | Set-Content "$root/$name/internal/handlers/$name.go"

@"
package service

import "$name/internal/repository"

func GetAll() map[string]interface{} {
	return map[string]interface{}{
		"service": "$name",
		"data": repository.FindAll(),
	}
}
"@ | Set-Content "$root/$name/internal/service/$name.go"

@"
package repository

func FindAll() []string {
	return []string{$items}
}
"@ | Set-Content "$root/$name/internal/repository/memory.go"

@"
package models

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
"@ | Set-Content "$root/$name/internal/models/item.go"
}

Create-Backend "users" "users" '"Mario", "Luigi"'
Create-Backend "orders" "orders" '"order-1", "order-2"'
Create-Backend "cart" "cart" '"item-1", "item-2"'

@"

```bash
cd users
go run ./cmd/users
