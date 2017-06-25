package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Index is a "placeholder" for the time being, until I work on the API
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome")
}

func main() {
	r := httprouter.New()

	r.GET("/", Index)

	http.ListenAndServe(":3000", r)
}
