package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/envs", ConfigMap)
	http.HandleFunc("/secret", Secret)
	http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Full Cycle \n"))
	user := os.Getenv("USER-CONTEXT")
	date := os.Getenv("DATE")
	fmt.Fprintf(w, "User-Context: %s | %s", user, date)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("go/envs/envs.txt")
	if err != nil {
		log.Fatalf("Error until read file", err)
	}

	fmt.Fprintf(w, "Envs: %s", string(data))

}

func Secret(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	fmt.Fprintf(w, "User: %s \nPassword: %s", user, password)
}
