package main

import (
	"net/http"

	. "github.com/rizkyalviandra/go-simple-api/contributor"
)

func main() {
	http.HandleFunc("/", SampleContributor)
	http.ListenAndServe(":4000", nil)
}
