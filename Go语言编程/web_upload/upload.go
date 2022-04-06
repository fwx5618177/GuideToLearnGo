package main

import (
	"io"
	"log"
	"net/http"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		io.WriteString(w, "Uploadeds.")

		return
	}
}

func main() {
	http.HandleFunc("/upload", uploadHandler)

	err := http.ListenAndServe(":8080", nil)

	if err == nil {
		log.Fatal("Listen and server: ", err.Error())
	}
}
