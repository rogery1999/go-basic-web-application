package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		const author = `Rogery`
		n, err := fmt.Fprintf(w, "Hello world!, %v", author)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Number of bytes written: %d\n", n)
	})

	_ = http.ListenAndServe("localhost:8080", nil)
}
