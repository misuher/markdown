package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static/"))) //working page
	http.HandleFunc("/preview", preview)                     //ajax url
	http.ListenAndServe(":3000", nil)
}

func preview(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}

	field := r.FormValue("textfield")
	//TODO: markdown parser call
	log.Println("Ajax:", field)
	w.Write([]byte(field))
}
