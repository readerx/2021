package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("%s <port> <name>\n", os.Args[0])
		return
	}

	port := os.Args[1]
	name := os.Args[2]

	http.DefaultServeMux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("requst from %s\n", request.RemoteAddr)
		writer.Write([]byte(name))
	})

	fmt.Printf("Listen at %s\n", port)
	http.ListenAndServe(":"+port, http.DefaultServeMux)
}
