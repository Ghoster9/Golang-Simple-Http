package handler

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

// root access
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// request halaman root
	log.Println(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	//jooin file
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	// error testing index html
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	// data static untuk read di html untuk interface bisa mereturnk bebas tipe data
	data := map[string]interface{}{
		"title":   "i'm learn golang web",
		"content": "I'm learning golang web with zaky",
	}
	// error testing res req
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

	// fmt.Fprintf(w, "Product page: %d", idNumb)

	data := map[string]interface{}{
		"content": idNumb,
	}

	// template can use for call page
	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	// error testing index html
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	// error testing res req
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
}
