package contributor

import (
	"fmt"
	"net/http"
)

func SampleContributor(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome contributor!")
}
