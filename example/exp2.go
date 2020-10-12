package main

import (
	"fmt"
	"log"
	"net/http"
)

type myServer struct {
}

func (*myServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!\n")
	fmt.Fprintf(w, "request data :\n %v\n", r)
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", &myServer{}))
}
