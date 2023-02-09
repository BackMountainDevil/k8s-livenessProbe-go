package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	started := time.Now()
	http.HandleFunc("/started", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		data := (time.Since(started)).String()
		w.Write([]byte(data))
	})
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		duration := time.Since(started)
		if duration.Seconds() > 40 {
			log.Println("receive health check. 500")
			w.WriteHeader(500)
			w.Write([]byte(fmt.Sprintf("error: %v", duration.Seconds())))
		} else {
			log.Println("receive health check. 200")
			w.WriteHeader(200)
			w.Write([]byte("ok"))
		}
	})
	http.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		log.Println("receive ready check. 200")
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	})
	log.Println("start server at localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
