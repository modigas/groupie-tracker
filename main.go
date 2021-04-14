package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("./public"))
	mux.Handle("/public/", http.StripPrefix("/public/", files))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/artist", artistInfo)
	server := http.Server{
		Addr:    ":9000",
		Handler: mux,
	}
	fmt.Println("Listening on port ", server.Addr)
	server.ListenAndServe()
}
