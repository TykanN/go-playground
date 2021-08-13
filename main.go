package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, this is tykan's go server")
	})

	fmt.Println("start server")

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
