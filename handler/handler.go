package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request)  {
	log.Println(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("welcome to home"))
}

func HelloHandler(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("hello world, saya sedang belajar"))
}

func MarioHandler(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("hello mario, sedang apa?"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request)  {
	id := r.URL.Query().Get("id")

	idNumb, err := strconv.Atoi(id)
	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	//w.Write([]byte("Product Page"))

	fmt.Fprintf(w, "Product page: %d", idNumb)
}
