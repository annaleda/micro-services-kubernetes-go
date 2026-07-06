package main

import (
"log"
"net/http"

"cart/internal/handlers"
)

func main() {
http.HandleFunc("/cart", handlers.GetAll)

log.Println("cart service running on :8080")
log.Fatal(http.ListenAndServe(":8080", nil))
}
