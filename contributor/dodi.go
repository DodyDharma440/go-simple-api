package contributor

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Book struct {
	Name        string `json:"name"`
	Page        int    `json:"page"`
	ReleaseYear int    `json:"release_year"`
}

func GetBookLists(w http.ResponseWriter, r *http.Request) {
	books := []Book{
		{Name: "Hello", Page: 100, ReleaseYear: 2002},
		{Name: "World", Page: 150, ReleaseYear: 2021},
	}

	res := map[string]interface{}{
		"data": books,
	}

	booksJson, err := json.Marshal(res)

	if err != nil {
		panic(err.Error())
	}

	w.Write([]byte(booksJson))
}

func DodiContributor(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome contributor!")
}
