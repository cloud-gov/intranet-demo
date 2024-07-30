package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hello from cloud.gov!</h1>\n<p>X-Forwarded-For: %s</p>\n", r.Header.Get("X-Forwarded-For"))
	})

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
