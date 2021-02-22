package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(respWriter http.ResponseWriter, req *http.Request) {
	fmt.Printf("Req uri: %s\n", req.URL.Path)
	code, err := respWriter.Write([]byte("Done and done"))
	if err != nil {
		fmt.Printf("%d: %v", code, err)
		respWriter.WriteHeader(500)
	}
}

func makeHandler(handler func(http.ResponseWriter, *http.Request) error, w http.ResponseWriter, r *http.Request) func(w2 http.ResponseWriter, r2 http.Request){
	 var err = handler(w, r)
	 if (err != nil) {

	 }
}

func helloHandler(respWriter http.ResponseWriter, req *http.Request) {
	fmt.Printf("Req uri: %s\n", req.URL.Path)
	code, err := respWriter.Write([]byte("Hello!!!!!!"))
	if err != nil {
		fmt.Printf("%d: %v", code, err)
		respWriter.WriteHeader(500)
	}
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}