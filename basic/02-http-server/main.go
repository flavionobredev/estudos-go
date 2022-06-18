package main

import (
	"fmt"
	"net/http"
	"os"
)

// to run: PORT=3030 go run main.go
func main() {
	port := fmt.Sprintf(":%v", os.Getenv("PORT"))
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Printf("[request]\t%v\n", request.URL.Path)
		response.Header().Add("Content-Type", "text/plain")
		response.Write([]byte("Hello, World!"))
	})
	http.ListenAndServe(port, nil)
}
