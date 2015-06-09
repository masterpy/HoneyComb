package router

import (
	mydb "Honeycomb/mysqlUtility"
	"fmt"
	"net/http"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)

	rows, _ := mydb.DBConn.Query("select * from users")

	for rows.Next() {
		var userid int
		var username string
		var useremail string
		_ = rows.Scan(&userid, &username, &useremail)

		fmt.Println(userid)
		fmt.Println(username)
		fmt.Println(useremail)
	}

	fmt.Fprintf(w, "Hello world.")
}

func LoadUserRouter() {
	http.HandleFunc("/user", sayhelloName)
}
