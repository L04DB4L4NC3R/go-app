package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	if port == "" {
		port = "3000"
	}
	fmt.Println("Listening on port " + port)
	http.ListenAndServe("0.0.0.0:"+port, nil)
}
