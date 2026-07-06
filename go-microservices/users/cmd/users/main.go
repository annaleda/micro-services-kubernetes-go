package main

import (
"log"
"net/http"

"users/internal/handlers"
)

func main() {
http.HandleFunc("/users", handlers.GetAll)

log.Println("users service running on :8080")
log.Fatal(http.ListenAndServe(":8080", nil))
}
