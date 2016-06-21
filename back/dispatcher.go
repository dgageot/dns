package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", count)
	http.ListenAndServe(":8080", nil)
}

func count(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello"))
	if err != nil {
		log.Println("Error", err)
		http.Error(w, err.Error(), 500)
		return
	}
}
