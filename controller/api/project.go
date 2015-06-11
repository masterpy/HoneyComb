package api

import (
	"fmt"
	"net/http"
)

func Projects(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Form)
	fmt.Fprintf(w, "123456")
}
