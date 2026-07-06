package main

import (
"log"
"net/http"

"orders/internal/handlers"
)

func main() {
http.HandleFunc("/orders", handlers.GetAll)

log.Println("orders service running on :8080")
log.Fatal(http.ListenAndServe(":8080", nil))
}
