package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func main() {
	PORT := 5001
	http.HandleFunc("/", helloHandler)
	log.Println(fmt.Sprintf("Listing for requests at http://localhost:%d", PORT))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil))
}
