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

type Project struct {
	ProjectCode   string `json:"project_code"`
	ProjectType   string `json:"project_type"`
	ProjectName   string `json:"project_name"`
	ProjectDetail string `json:"project_detail"`
	Manager       string `json:"manager"`
	StartDate     string `json:"start_date"`
	DueDate       string `json:"due_date"`
	RealStartDate string `json:"real_start_date"`
	RealDueDate   string `json:"real_due_date"`
	Status        int    `json:"status"`
	WinPath       string `json:"win_path"`
	LinuxPath     string `json:"linux_path"`
	MacOSPath     string `json:"macos_path"`
}

func AddProject(w http.ResponseWriter, r *http.Request) {
	project_name := strings.Join(r.Form["project_name"], "")
	project_detail := strings.Join(r.Form["project_detail"], "")
	project_code := ut.GenerateCode(project_name)

	fmt.Println("Add project", project_name)

	stmt, err := mydb.DBConn.Prepare("INSERT INTO project(project_code, project_name, project_detail) VALUES(?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(project_code, project_name, project_detail)
	if err != nil {
		panic(err.Error())
	}

	proj := QueryProjectByCode(project_code)
	res, _ := json.Marshal(proj)

	fmt.Fprintf(w, string(res))
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	project_code := strings.Join(r.Form["project_code"], "")
	project_name := strings.Join(r.Form["project_name"], "")
	project_type := strings.Join(r.Form["project_type"], "")
	project_detail := strings.Join(r.Form["project_detail"], "")
	manager := strings.Join(r.Form["manager"], "")
	start_date := strings.Join(r.Form["start_date"], "")
	due_date := strings.Join(r.Form["due_date"], "")
	status_str := strings.Join(r.Form["status"], "")
	status, _ := strconv.Atoi(status_str)
	win_path := strings.Join(r.Form["win_path"], "")
	linux_path := strings.Join(r.Form["linux_path"], "")
	macos_path := strings.Join(r.Form["macos_path"], "")

	fmt.Println("Update project", project_name)

	stmt, err := mydb.DBConn.Prepare("UPDATE project SET project_name = ?, project_type = ?, project_detail = ?, manager = ?, start_date = ?, due_date = ?, status = ?, win_path = ?, linux_path = ?, macos_path = ? WHERE project_code = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(project_name, project_type, project_detail, manager, start_date, due_date, status, win_path, linux_path, macos_path, project_code)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "[true,]")
}

func RemoveProject(w http.ResponseWriter, r *http.Request) {
	project_name := strings.Join(r.Form["project_name"], "")
	project_code := strings.Join(r.Form["project_code"], "")

	fmt.Println("Remove project", project_name)

	// 删除所有的任务
	RemoveProjectMissions(project_code)

	// 删除项目
	stmt, err := mydb.DBConn.Prepare("DELETE FROM project WHERE project_code=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(project_code)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "[true,]")
}

func QueryProjectByCode(projectCode string) (project Project) {
	stmt, err := mydb.DBConn.Prepare("SELECT project_code, project_type, project_name, project_detail, manager, start_date, due_date, status, win_path, linux_path, macos_path FROM project WHERE project_code=?")
	defer stmt.Close()
	if err != nil {
		panic(err.Error())
	}

	result, err := stmt.Query(projectCode)
	defer result.Close()
	if err != nil {
		panic(err.Error())
	}

	//var project Project
	if result.Next() {
		err = result.Scan(&(project.ProjectCode), &(project.ProjectCode), &(project.ProjectName), &(project.ProjectDetail), &(project.Manager), &(project.StartDate), &(project.DueDate), &(project.Status), &(project.WinPath), &(project.LinuxPath), &(project.MacOSPath))
		if err != nil {
			panic(err.Error())
		}
	}
	return project
}

func GetProjects(w http.ResponseWriter, r *http.Request) {
	result, err := mydb.DBConn.Query("SELECT project_code, project_type, project_name, project_detail, manager, start_date, due_date, status, win_path, linux_path, macos_path FROM project")
	defer result.Close()
	if err != nil {
		panic(err.Error())
	}

	var projs []interface{}

	for result.Next() {
		var project Project
		err = result.Scan(&(project.ProjectCode), &(project.ProjectType), &(project.ProjectName), &(project.ProjectDetail), &(project.Manager), &(project.StartDate), &(project.DueDate), &(project.Status), &(project.WinPath), &(project.LinuxPath), &(project.MacOSPath))
		if err != nil {
			panic(err.Error())
		}
		projs = append(projs, project)
	}

	res, _ := json.Marshal(projs)

	fmt.Fprintf(w, string(res))
}
