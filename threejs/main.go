package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "res/index.html") })
	http.HandleFunc("/main.js",
		func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "res/main.js") })
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe fatal error:", err.Error())
	}
}
