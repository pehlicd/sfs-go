package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./files/"))

	mux.HandleFunc("/", rootHandler)

	mux.Handle("/files/", http.StripPrefix("/files", fileServer))

	log.Println("Starting server on :4000")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	err := http.ListenAndServe(":4000", logRequest(mux))
	if err != nil {
		log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to file server</h1><div>My mission is serving files for you.</div>")
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
