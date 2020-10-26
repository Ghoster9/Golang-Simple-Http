package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

// root access
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"title":   "i'm learn golang web",
		"Content": "I'm learning golang web with zaky",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
}

// /hello
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world, saya sedang belajar"))
}

// /mario
func MarioHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello mario, sedang apa?"))
}

// /product?id=1
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idNumb, err := strconv.Atoi(id)
	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	//w.Write([]byte("Product Page"))

	fmt.Fprintf(w, "Product page: %d", idNumb)
}
