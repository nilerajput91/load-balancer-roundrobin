package main

import (
	"fmt"
	"net/http"
)

var (
	port = "8082"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":"+port, nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from port %s!", port)
}
