package api

import (
	mydb "Honeycomb/mysqlUtility"
	ut "Honeycomb/utility"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Mission struct {
	MissionCode   string `json:"mission_code"`
	MissionName   string `json:"mission_name"`
	MissionType   string `json:"mission_type"`
	MissionDetail string `json:"mission_detail"`
	ProjectCode   string `json:"project_code"`
	AssignTo      string `json:"assignto"`
	HasChild      string `json:"has_child"`
	ParentCode    string `json:"parent_code"`
	ChildIndex    int    `json:"child_index"`
	Status        int    `json:"status"`
}

// CREATE TABLE `mission` (
//     `mission_id` int unsigned NOT NULL AUTO_INCREMENT,
//     `mission_code` char(32) NOT NULL unique,
//     `mission_name` char(50) NOT NULL,                      #任务的名称
//     `mission_type` char(32) NOT NULL,                      #存储mission_type_code，任务的类型
//     `mission_detail` varchar(200) NOT NULL,                #给任务一段更详细的文本说明
//     `project_code` char(32) NOT NULL,                      #任务所属project
//     `has_child` tinyint default 0 NOT NULL,                #是否包含子任务，0不是，1是
//     `parent_code` char(32) NOT NULL,                       #父级mission_code
//     `child_index` int default 0 NOT NULL,                  #作为子任务，其所在父下的顺序
//     `plan_begin_datetime` datetime,                        #计划开始时间
//     `plan_end_datetime` datetime,                          #计划结束时间
//     `real_begin_datetime` datetime,                        #实际开始时间
//     `real_end_datetime` datetime,                          #实际结束时间
//     `assingto` char(32),                                   #存储`user_code`，任务的负责人
//     `status` int default 0 NOT NULL,                       #0指定了人但未开始，1已经完成,2已经通过，3进行中，4未指定人
//     `picture` mediumtext,                                  #照片的base64编码
//     `insert_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//     `update_datetime` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
//     PRIMARY KEY (`mission_id`),
//     INDEX(`mission_code`),
//     INDEX(`project_code`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

func AddMission(w http.ResponseWriter, r *http.Request) {
	mission_name := strings.Join(r.Form["mission_name"], "")
	mission_type := strings.Join(r.Form["mission_type"], "")
	mission_detail := ""
	project_code := strings.Join(r.Form["project_code"], "")
	has_child := 1
	parent_code := strings.Join(r.Form["parent_code"], "")
	child_index := 0
	status := 0

	mission_code := ut.GenerateCode(mission_name)

	fmt.Println("Add mission", mission_name)

	stmt, err := mydb.DBConn.Prepare("INSERT INTO mission(mission_code, mission_name, mission_type, mission_detail, project_code, has_child, parent_code, child_index, status) VALUES(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(mission_code, mission_name, mission_type, mission_detail, project_code, has_child, parent_code, child_index, status)
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

func QueryMissionByCode(missionCode string) (mission Mission) {
	stmt, err := mydb.DBConn.Prepare("SELECT mission_code, mission_name, mission_type, mission_detail, project_code, has_child, parent_code, child_index, status FROM mission WHERE mission_code=?")
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
		err = result.Scan(&(mission.MissionCode), &(mission.MissionName), &(mission.MissionType), &(mission.MissionDetail), &(mission.ProjectCode), &(mission.HasChild), &(mission.ParentCode), &(mission.ChildIndex), &(mission.Status))
		if err != nil {
			panic(err.Error())
		}
	}
	return mission
}

func GetMissions(w http.ResponseWriter, r *http.Request) {
	project_code := strings.Join(r.Form["project_code"], "")

	stmt, err := mydb.DBConn.Prepare("SELECT mission_code, mission_name, mission_type, mission_detail, project_code, has_child, parent_code, child_index, status FROM mission WHERE project_code=?")
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
		err = result.Scan(&(mission.MissionCode), &(mission.MissionName), &(mission.MissionType), &(mission.MissionDetail), &(mission.ProjectCode), &(mission.HasChild), &(mission.ParentCode), &(mission.ChildIndex), &(mission.Status))
		if err != nil {
			panic(err.Error())
		}

		missions = append(missions, mission)
	}

	res, _ := json.Marshal(missions)

	fmt.Fprintf(w, string(res))
}
