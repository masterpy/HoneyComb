package api

import (
	mydb "Honeycomb/mysqlUtility"
	ut "Honeycomb/utility"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	//"time"
)

type Mission struct {
	MissionCode        string `json:"mission_code"`
	MissionName        string `json:"mission_name"`
	MissionType        string `json:"mission_type"`
	Priority           string `json:"priority"`
	DegreeOfDifficulty string `json:"degree_of_difficulty"`
	ProjectCode        string `json:"project_code"`
	HasChild           int    `json:"has_child"`
	ParentCode         string `json:"parent_code"`
	ChildIndex         int    `json:"child_index"`
	StartDate          string `json:"start_date"`
	DueDate            string `json:"due_date"`
	RealStartDate      string `json:"real_start_date"`
	RealDueDate        string `json:"real_due_date"`
	AssignTo           string `json:"assign_to"`
	Status             int    `json:"status"`
}

func AddMission(w http.ResponseWriter, r *http.Request) {
	mission_name := strings.Join(r.Form["mission_name"], "")
	mission_type := strings.Join(r.Form["mission_type"], "")
	project_code := strings.Join(r.Form["project_code"], "")
	parent_code := strings.Join(r.Form["parent_code"], "")
	has_child := 1
	child_index := 0
	status := 0

	mission_code := ut.GenerateCode(mission_name)

	fmt.Println("Add mission", mission_name)

	stmt, err := mydb.DBConn.Prepare("INSERT INTO mission(mission_code, mission_name, mission_type, project_code, has_child, parent_code, child_index, status) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(mission_code, mission_name, mission_type, project_code, has_child, parent_code, child_index, status)
	if err != nil {
		panic(err.Error())
	}

	mission := QueryMissionByCode(mission_code)
	res, _ := json.Marshal(mission)

	fmt.Fprintf(w, string(res))
}

// 现阶段直接移除，不做任何判断，以后要根据任务的状态以及分配情况选择直接删除还是标记删除
func RemoveMission(w http.ResponseWriter, r *http.Request) {
	mission_name := strings.Join(r.Form["mission_name"], "")
	mission_code := strings.Join(r.Form["mission_code"], "")

	fmt.Println("Remove mission", mission_name)

	stmt, err := mydb.DBConn.Prepare("DELETE FROM mission WHERE mission_code=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(mission_code)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "[true,]")
}

func UpdateMission(w http.ResponseWriter, r *http.Request) {
	mission_code := strings.Join(r.Form["mission_code"], "")
	mission_name := strings.Join(r.Form["mission_name"], "")
	mission_type := strings.Join(r.Form["mission_type"], "")
	assign_to := strings.Join(r.Form["assign_to"], "")
	status_str := strings.Join(r.Form["status"], "")
	status, _ := strconv.Atoi(status_str)
	start_date := strings.Join(r.Form["start_date"], "")
	due_date := strings.Join(r.Form["due_date"], "")

	fmt.Println("Update mission", mission_name)

	stmt, err := mydb.DBConn.Prepare("UPDATE mission SET mission_name = ?, mission_type = ?, assign_to = ?, status = ?, start_date = ?, due_date = ? WHERE mission_code = ?")

	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(mission_name, mission_type, assign_to, status, start_date, due_date, mission_code)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "[true,]")
}

func QueryMissionByCode(missionCode string) (mission Mission) {
	stmt, err := mydb.DBConn.Prepare("SELECT mission_code, mission_name, mission_type, project_code, has_child, parent_code, child_index, status, start_date, due_date FROM mission WHERE mission_code=?")
	defer stmt.Close()
	if err != nil {
		panic(err.Error())
	}

	result, err := stmt.Query(missionCode)
	defer result.Close()
	if err != nil {
		panic(err.Error())
	}

	//var project Project
	if result.Next() {
		err = result.Scan(&(mission.MissionCode), &(mission.MissionName), &(mission.MissionType), &(mission.ProjectCode), &(mission.HasChild), &(mission.ParentCode), &(mission.ChildIndex), &(mission.Status), &(mission.StartDate), &(mission.DueDate))
		if err != nil {
			panic(err.Error())
		}
	}
	return mission
}

func GetMissions(w http.ResponseWriter, r *http.Request) {
	project_code := strings.Join(r.Form["project_code"], "")

	stmt, err := mydb.DBConn.Prepare("SELECT mission_code, mission_name, mission_type, project_code, has_child, parent_code, child_index, assign_to, status, start_date, due_date FROM mission WHERE project_code=?")
	defer stmt.Close()
	if err != nil {
		panic(err.Error())
	}

	result, err := stmt.Query(project_code)
	defer result.Close()
	if err != nil {
		panic(err.Error())
	}

	var missions []interface{}

	for result.Next() {
		var mission Mission
		err = result.Scan(&(mission.MissionCode), &(mission.MissionName), &(mission.MissionType), &(mission.ProjectCode), &(mission.HasChild), &(mission.ParentCode), &(mission.ChildIndex), &(mission.AssignTo), &(mission.Status), &(mission.StartDate), &(mission.DueDate))
		if err != nil {
			panic(err.Error())
		}

		missions = append(missions, mission)
	}

	res, _ := json.Marshal(missions)

	fmt.Fprintf(w, string(res))
}
