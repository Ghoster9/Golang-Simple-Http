package main

import (
	"golangweb/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	//aboutHandler := func(w http.ResponseWriter, r *http.Request){
	//	w.Write([]byte("About page"))
	//}

	//root
	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/mario", handler.MarioHandler)
	mux.HandleFunc("/product", handler.ProductHandler)
	mux.HandleFunc("/post-get", handler.PostGet)

	// ini untuk membaca file asset yang bersi style
	fileServer := http.FileServer(http.Dir("assets"))

	//mengload dengan rooting
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	//mux.HandleFunc("/about", aboutHandler)
	//mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("Profile"))
	//})

	log.Println("Starting web on port : 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
