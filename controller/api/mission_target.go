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

type MissionTarget struct {
	TargetCode  string `json:"target_code"`
	MissionCode string `json:"mission_code"`
	TargetName  string `json:"target_name"`
	TargetType  int    `json:"target_type"`
	Comment     string `json:"comment"`
}

func AddTarget(w http.ResponseWriter, r *http.Request) {
	mission_code := strings.Join(r.Form["mission_code"], "")
	target_name := strings.Join(r.Form["target_name"], "")
	type_str := strings.Join(r.Form["target_type"], "")
	target_type, _ := strconv.Atoi(type_str)
	comment := strings.Join(r.Form["comment"], "")

	target_code := ut.GenerateCode(target_name + comment)

	fmt.Println("Add Require for mission ", target_name)

	stmt, err := mydb.DBConn.Prepare("INSERT INTO mission_target(target_code, mission_code, target_name, target_type, comment) VALUES(?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(target_code, mission_code, target_name, target_type, comment)
	if err != nil {
		panic(err.Error())
	}

	target := QueryTargetByCode(target_code)
	res, _ := json.Marshal(target)

	fmt.Fprintf(w, string(res))
}

func UpdateTarget(w http.ResponseWriter, r *http.Request) {
	target_code := strings.Join(r.Form["target_code"], "")
	target_name := strings.Join(r.Form["target_name"], "")
	type_str := strings.Join(r.Form["target_type"], "")
	target_type, _ := strconv.Atoi(type_str)
	comment := strings.Join(r.Form["comment"], "")

	fmt.Println("Update mission require", target_code)

	stmt, err := mydb.DBConn.Prepare("UPDATE mission_target SET target_name=?, target_type=?, comment = ? WHERE target_code = ?")

	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(target_name, target_type, comment, target_code)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "[true,]")
}

// 现阶段直接移除，不做任何判断，以后要根据任务的状态以及分配情况选择直接删除还是标记删除
func RemoveTarget(w http.ResponseWriter, r *http.Request) {
	target_code := strings.Join(r.Form["target_code"], "")

	fmt.Println("Remove mission target.", target_code)

	stmt, err := mydb.DBConn.Prepare("DELETE FROM mission_target WHERE target_code=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(target_code)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "[true,]")
}

func QueryTargetByCode(target_code string) (target MissionTarget) {
	stmt, err := mydb.DBConn.Prepare("SELECT target_code, mission_code, target_name, target_type, comment FROM mission_target WHERE target_code=?")
	defer stmt.Close()
	if err != nil {
		panic(err.Error())
	}

	result, err := stmt.Query(target_code)
	defer result.Close()
	if err != nil {
		panic(err.Error())
	}

	//var project Project
	if result.Next() {
		err = result.Scan(&(target.TargetCode), &(target.MissionCode), &(target.TargetName), &(target.TargetType), &(target.Comment))
		if err != nil {
			panic(err.Error())
		}
	}
	return target
}

func GetMissionTarget(w http.ResponseWriter, r *http.Request) {
	mission_code := strings.Join(r.Form["mission_code"], "")

	stmt, err := mydb.DBConn.Prepare("SELECT target_code, mission_code, target_name, target_type, comment FROM mission_target WHERE mission_code=?")
	defer stmt.Close()
	if err != nil {
		panic(err.Error())
	}

	result, err := stmt.Query(mission_code)
	defer result.Close()
	if err != nil {
		panic(err.Error())
	}

	var target_list []interface{}

	for result.Next() {
		var target MissionTarget
		err = result.Scan(&(target.TargetCode), &(target.MissionCode), &(target.TargetName), &(target.TargetType), &(target.Comment))
		if err != nil {
			panic(err.Error())
		}

		target_list = append(target_list, target)
	}

	res, _ := json.Marshal(target_list)

	fmt.Fprintf(w, string(res))
}
