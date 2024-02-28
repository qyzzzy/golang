package main

import (
	"fmt"
	"net/http"
)

func handlers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>hello ,world! %s </h1>", r.URL)
}
func main() {
	http.HandleFunc("/", handlers)
	http.ListenAndServe(":8080", nil)
}
