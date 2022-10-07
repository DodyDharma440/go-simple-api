package main

import (
	"fmt"
	"log"
	"net/http"

	. "github.com/rizkyalviandra/go-simple-api/contributor"
)

func main() {
	http.HandleFunc("/", SampleContributor)
	http.HandleFunc("/dodi", DodiContributor)
	http.HandleFunc("/books", GetBookLists)

	fmt.Println("App is running on http://localhost:4000")
	if err := http.ListenAndServe(":4000", nil); err != nil {
		log.Fatal(err)
	}
}
