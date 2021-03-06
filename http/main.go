package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz/", healthzHandler)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("IP is " + r.RemoteAddr + " HTTP Code is " + http.StatusText(200))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	for idx, value := range r.Header {
		w.Header().Add(idx, strings.Join(value, "-"))
	}
	w.Header().Add("VERSION", os.Getenv("VERSION"))
	log.Println("IP is " + r.RemoteAddr + " HTTP Code is " + http.StatusText(200))
}
