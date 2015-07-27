package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/misuher/markdown/markparser"
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
	mark := markdown.NewParser(strings.NewReader(field))
	log.Println("Ajax:", field)
	w.Write([]byte(mark.Markdown()))
}
