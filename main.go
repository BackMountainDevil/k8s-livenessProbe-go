package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("start server at localhost:8081")
	http.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.ParseForm)
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	})
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
