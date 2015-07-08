package api

import (
	mydb "Honeycomb/mysqlUtility"
	ut "Honeycomb/utility"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type MissionRequire struct {
	RequireCode string `json:"require_code"`
	MissionCode string `json:"mission_code"`
	RequireType int    `json:"require_type"`
	Content     string `json:"content"`
}

func AddRequire(w http.ResponseWriter, r *http.Request) {
	mission_code := strings.Join(r.Form["mission_code"], "")
	mission_name := strings.Join(r.Form["mission_name"], "")
	type_str := strings.Join(r.Form["require_type"], "")
	require_type, _ := strconv.Atoi(type_str)
	content := strings.Join(r.Form["content"], "")

	require_code := ut.GenerateCode(content)

	fmt.Println("Add Require for mission ", mission_name)

	stmt, err := mydb.DBConn.Prepare("INSERT INTO mission_require(require_code, mission_code, require_type, content) VALUES(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(require_code, mission_code, require_type, content)
	if err != nil {
		panic(err.Error())
	}

	require := QueryRequireByCode(require_code)
	res, _ := json.Marshal(require)

	fmt.Fprintf(w, string(res))
}

func UpdateRequire(w http.ResponseWriter, r *http.Request) {
	require_code := strings.Join(r.Form["require_code"], "")
	// mission_code := strings.Join(r.Form["mission_code"], "")
	// type_str := strings.Join(r.Form["require_type"], "")
	// require_type, _ := strconv.Atoi(type_str)
	content := strings.Join(r.Form["content"], "")

	fmt.Println("Update mission require", require_code)

	stmt, err := mydb.DBConn.Prepare("UPDATE mission_require SET content = ? WHERE require_code = ?")

	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(content, require_code)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "[true,]")
}

// 现阶段直接移除，不做任何判断，以后要根据任务的状态以及分配情况选择直接删除还是标记删除
func RemoveRequire(w http.ResponseWriter, r *http.Request) {
	require_code := strings.Join(r.Form["require_code"], "")

	fmt.Println("Remove mission require.", require_code)

	stmt, err := mydb.DBConn.Prepare("DELETE FROM mission_require WHERE require_code=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(require_code)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "[true,]")
}

func QueryRequireByCode(require_code string) (require MissionRequire) {
	stmt, err := mydb.DBConn.Prepare("SELECT require_code, mission_code, require_type, content FROM mission_require WHERE require_code=?")
	defer stmt.Close()
	if err != nil {
		panic(err.Error())
	}

	result, err := stmt.Query(require_code)
	defer result.Close()
	if err != nil {
		panic(err.Error())
	}

	//var project Project
	if result.Next() {
		err = result.Scan(&(require.RequireCode), &(require.MissionCode), &(require.RequireType), &(require.Content))
		if err != nil {
			panic(err.Error())
		}
	}
	return require
}

func GetMissionRequire(w http.ResponseWriter, r *http.Request) {
	mission_code := strings.Join(r.Form["mission_code"], "")

	stmt, err := mydb.DBConn.Prepare("SELECT require_code, mission_code, require_type, content FROM mission_require WHERE mission_code=?")
	defer stmt.Close()
	if err != nil {
		panic(err.Error())
	}

	result, err := stmt.Query(mission_code)
	defer result.Close()
	if err != nil {
		panic(err.Error())
	}

	var require_list []interface{}

	for result.Next() {
		var require MissionRequire
		err = result.Scan(&(require.RequireCode), &(require.MissionCode), &(require.RequireType), &(require.Content))
		if err != nil {
			panic(err.Error())
		}

		require_list = append(require_list, require)
	}

	res, _ := json.Marshal(require_list)

	fmt.Fprintf(w, string(res))
}
