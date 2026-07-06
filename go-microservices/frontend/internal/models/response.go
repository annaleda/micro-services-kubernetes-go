package models

type BackendResponse struct {
Service string   json:"service"
Data    []string json:"data"
}

type PageData struct {
Users  BackendResponse
Orders BackendResponse
Cart   BackendResponse
}
