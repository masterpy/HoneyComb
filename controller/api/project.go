package api

import (
	mydb "Honeycomb/mysqlUtility"
	ut "Honeycomb/utility"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Project struct {
	ProjectCode    string `json:"project_code"`
	ProjectName    string `json:"project_name"`
	ProjectDetail  string `json:"project_detail"`
	PersonInCharge string `json:"person_in_charge"`
	CompanyCode    string `json:"company_code"`
	Status         int    `json:"status"`
}

func AddProject(w http.ResponseWriter, r *http.Request) {
	project_name := strings.Join(r.Form["project_name"], "")
	project_detial := strings.Join(r.Form["project_detial"], "")
	project_code := ut.GenerateCode(project_name)

	fmt.Println("Add project", project_name)

	stmt, err := mydb.DBConn.Prepare("INSERT INTO project(project_code, project_name, project_detail, person_in_charge, company_code, status) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(project_code, project_name, project_detial, "", "", 0)
	if err != nil {
		panic(err.Error())
	}

	proj := QueryProjectByCode(project_code)
	res, _ := json.Marshal(proj)

	fmt.Fprintf(w, string(res))
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	project_name := strings.Join(r.Form["project_name"], "")
	project_detial := strings.Join(r.Form["project_detial"], "")
	project_code := ut.GenerateCode(project_name)

	fmt.Println("Add project", project_name)

	stmt, err := mydb.DBConn.Prepare("INSERT INTO project(project_code, project_name, project_detail, person_in_charge, company_code, status) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(project_code, project_name, project_detial, "", "", 0)
	if err != nil {
		panic(err.Error())
	}

	proj := QueryProjectByCode(project_code)
	res, _ := json.Marshal(proj)

	fmt.Fprintf(w, string(res))
}

func RemoveProject(w http.ResponseWriter, r *http.Request) {
	project_name := strings.Join(r.Form["project_name"], "")
	project_detial := strings.Join(r.Form["project_detial"], "")
	project_code := ut.GenerateCode(project_name)

	fmt.Println("Add project", project_name)

	stmt, err := mydb.DBConn.Prepare("INSERT INTO project(project_code, project_name, project_detail, person_in_charge, company_code, status) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(project_code, project_name, project_detial, "", "", 0)
	if err != nil {
		panic(err.Error())
	}

	proj := QueryProjectByCode(project_code)
	res, _ := json.Marshal(proj)

	fmt.Fprintf(w, string(res))
}

func QueryProjectByCode(projectCode string) (project Project) {
	stmt, err := mydb.DBConn.Prepare("SELECT project_code, project_name, project_detail, person_in_charge, company_code, status FROM project WHERE project_code=?")
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
		err = result.Scan(&(project.ProjectCode), &(project.ProjectName), &(project.ProjectDetail),
			&(project.PersonInCharge), &(project.CompanyCode),
			&(project.Status))
		if err != nil {
			panic(err.Error())
		}
	}
	return project
}

func GetProjects(w http.ResponseWriter, r *http.Request) {
	result, err := mydb.DBConn.Query("SELECT project_code, project_name, project_detail, person_in_charge, company_code, status FROM project")
	defer result.Close()
	if err != nil {
		panic(err.Error())
	}

	var projs []interface{}

	for result.Next() {
		var project Project
		err = result.Scan(&(project.ProjectCode), &(project.ProjectName), &(project.ProjectDetail),
			&(project.PersonInCharge), &(project.CompanyCode),
			&(project.Status))
		if err != nil {
			panic(err.Error())
		}
		projs = append(projs, project)
	}

	res, _ := json.Marshal(projs)

	fmt.Fprintf(w, string(res))
}
