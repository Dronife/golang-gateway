package main

import (
    "fmt"
    "net/http"
)

func StartHTTPServer() {
    http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello from API!")
    })
    fmt.Println("HTTP server listening on :8080 !!")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("HTTP server error:", err)
    }
}
