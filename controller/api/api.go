package api

import (
	"fmt"
	"net/http"
	"strings"
)

func checkAPIKey(key string) bool {
	if key == "123456" {
		return true
	}

	return false

}

func Hub(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	k := strings.Join(r.Form["api-key"], "")

	if !checkAPIKey(k) {
		fmt.Println("Api key is wrong.")
		return
	}

	fmt.Println(r.Form)

	fmt.Fprintf(w, "Hello world.")

	Projects(w, r)
}
