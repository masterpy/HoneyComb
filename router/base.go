package router

import (
	"net/http"
)

func LoadBaseRouter() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./static"))))
}
