package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is from the snippet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Server starting at :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
