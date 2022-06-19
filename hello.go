package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}
func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", hello)
	log.Println("Listening on portc:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
