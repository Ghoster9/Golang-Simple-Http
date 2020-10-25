package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/mario", marioHandler)


	log.Println("Starting web on port : 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func homeHandler(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("welcome to home"))
}

func helloHandler(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("hello world, saya sedang belajar"))
}

func marioHandler(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("hello mario, sedang apa?"))
}