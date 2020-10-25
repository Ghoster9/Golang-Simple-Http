package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()

	//aboutHandler := func(w http.ResponseWriter, r *http.Request){
	//	w.Write([]byte("About page"))
	//}

	//root
	mux.HandleFunc("/", homeHandler)

	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/mario", marioHandler)
	mux.HandleFunc("/product", productHandler)

	//mux.HandleFunc("/about", aboutHandler)
	//mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("Profile"))
	//})


	log.Println("Starting web on port : 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func homeHandler(w http.ResponseWriter, r *http.Request)  {
	log.Println(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("welcome to home"))
}

func helloHandler(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("hello world, saya sedang belajar"))
}

func marioHandler(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("hello mario, sedang apa?"))
}

func productHandler(w http.ResponseWriter, r *http.Request)  {
	id := r.URL.Query().Get("id")

	idNumb, err := strconv.Atoi(id)
	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	//w.Write([]byte("Product Page"))

	fmt.Fprintf(w, "Product page: %d", idNumb)
}