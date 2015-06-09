package main

import (
	_ "Honeycomb/mysqlUtility"
	"Honeycomb/router"
	"log"
	"net/http"
)

func main() {
	router.LoadBaseRouter()
	router.LoadUserRouter()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}

}
