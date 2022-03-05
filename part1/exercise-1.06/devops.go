pckage main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world!")
}

func main() {
	addr := ":5000"
	http.HandleFunc("/", hello)
	log.Println("Server started in port", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
