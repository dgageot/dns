package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", count)
	http.ListenAndServe(":8080", nil)
}

func count(w http.ResponseWriter, r *http.Request) {
	addrs, err := net.LookupHost("back")
	if err != nil {
		log.Println("Error", err)
		http.Error(w, err.Error(), 500)
		return
	}

	_, err = w.Write([]byte(fmt.Sprintf("I see %d java container(s)", len(addrs))))
	if err != nil {
		log.Println("Error", err)
		http.Error(w, err.Error(), 500)
		return
	}
}
