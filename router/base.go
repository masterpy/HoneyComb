package router

import (
	"Honeycomb/controller/api"
	"net/http"
)

func LoadBaseRouter() {
	// 设置静态资源
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/api", api.Hub)
}
