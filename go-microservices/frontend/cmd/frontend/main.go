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
