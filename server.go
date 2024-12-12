package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Full Cycle \n"))
	user := os.Getenv("USER-CONTEXT")
	date := os.Getenv("DATE")
	fmt.Fprintf(w, "User-Context: %s | %s", user, date)
}
