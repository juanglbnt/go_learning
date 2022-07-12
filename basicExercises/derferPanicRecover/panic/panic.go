package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from go"))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Here the panic:")
		panic(err.Error())
	}
}
