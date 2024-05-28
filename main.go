package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func healthCheck(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("I feel very healthy"))
}

func  main()  {
	http.HandleFunc("/health", healthCheck)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}