package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	for name, values := range r.Header {
		for _, value := range values {
			fmt.Fprintf(w, "%s: %s\n", name, value)
		}
	}

	fmt.Fprint(w, "\n\n")

	body, err := io.ReadAll(r.Body)
	if err == nil {
		w.Write(body)
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("listening on :80")
	http.ListenAndServe(":80", nil)
}
