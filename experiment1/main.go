package main

import (
	"log"
	"net/http"
)

func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middleWareOne")
		next.ServeHTTP(w, r)
		log.Println("Completed middleWareone")
	})
}

func middlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middleWareTwo")
		next.ServeHTTP(w, r)
		log.Println("Complted middleWareTwo")
	})
}

func final(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing final")
	w.Write([]byte("OKOKOK"))
	log.Println("Complted final")
}

func main() {
	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(final)
	mux.Handle("/", middlewareOne(middlewareTwo(finalHandler)))

	http.ListenAndServe(":8080", mux)
}
