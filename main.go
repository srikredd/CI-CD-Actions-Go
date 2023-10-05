package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", helloHandler)
    port := 8080
    addr := fmt.Sprintf(":%d", port)
    fmt.Printf("Server is listening on port %d...\n", port)
    err := http.ListenAndServe(addr, nil)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}
