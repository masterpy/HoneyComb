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

func defaultCmd(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "")
}

func Hub(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	k := strings.Join(r.Form["api-key"], "")

	if !checkAPIKey(k) {
		fmt.Println("Api key is wrong.")
		return
	}

	cmd := strings.Join(r.Form["command"], "")

	switch cmd {
	case "addProject":
		AddProject(w, r)
	case "getProjects":
		GetProjects(w, r)
	case "addMission":
		AddMission(w, r)
	case "getMissions":
		GetMissions(w, r)
	default:
		defaultCmd(w, r)
	}
}
