package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Hello world!</h1>")
}

func main() {
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
