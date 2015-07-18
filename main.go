package main

import "net/http"

func main() {
	http.HandleFunc("/", markdown)       //working page
	http.HandleFunc("/preview", preview) //ajax url
	http.ListenAndServe(":3000", nil)
}

func markdown(w http.ResponseWriter, r *http.Request) {

}

func preview(w http.ResponseWriter, r *http.Request) {

}
