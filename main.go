package main

import (
	"fmt"
	"net/http"
	"log"
)

const webContent = " Hudson Coutinho - DevOps Engineer  "

func main() {
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, webContent)
}
