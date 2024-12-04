package main

import (
    "fmt"
    "net/http"
    "os"
)

func helloWoldPage(w http.ResponseWriter, r *http.Request) {
    hostname, err := os.Hostname()
    if err != nil {
        http.Error(w, "Could not get hostname", http.StatusInternalServerError)
        return
    }
    fmt.Fprintf(w, "test 2 Hostname: %s", hostname)
}

func main() {
    http.HandleFunc("/", helloWoldPage)
    http.ListenAndServe(":8080", nil)
}