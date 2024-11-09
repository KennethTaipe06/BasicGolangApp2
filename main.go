package main

import (
	"fmt"
	"net/http"
)

func helloWoldPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It works 3/5 golang")
}

func main() {
	http.HandleFunc("/", helloWoldPage)
	http.ListenAndServe(":8080", nil)
}
