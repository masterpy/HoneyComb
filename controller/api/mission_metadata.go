package api

import (
	mydb "Honeycomb/mysqlUtility"
	ut "Honeycomb/utility"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type MissionMetadata struct {
	MetadataCode  string `json:"metadata_code"`
	MissionCode   string `json:"mission_code"`
	MetadataName  string `json:"metadata_name"`
	MetadataValue string `json:"metadata_value"`
}

func AddMetadata(w http.ResponseWriter, r *http.Request) {
	mission_code := strings.Join(r.Form["mission_code"], "")
	metadata_name := strings.Join(r.Form["metadata_name"], "")
	metadata_value := strings.Join(r.Form["metadata_value"], "")

	metadata_code := ut.GenerateCode(metadata_name + metadata_value)

	fmt.Println("Add Metadata for mission ", metadata_name)

	stmt, err := mydb.DBConn.Prepare("INSERT INTO mission_metadata(metadata_code, mission_code, metadata_name, metadata_value) VALUES(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(metadata_code, mission_code, metadata_name, metadata_value)
	if err != nil {
		panic(err.Error())
	}

	metadata := QueryMetadataByCode(metadata_code)
	res, _ := json.Marshal(metadata)

	fmt.Fprintf(w, string(res))
}

func UpdateMetadata(w http.ResponseWriter, r *http.Request) {
	metadata_code := strings.Join(r.Form["metadata_code"], "")
	metadata_name := strings.Join(r.Form["metadata_name"], "")
	metadata_value := strings.Join(r.Form["metadata_value"], "")

	fmt.Println("Update mission metadata", metadata_name)

	stmt, err := mydb.DBConn.Prepare("UPDATE mission_metadata SET metadata_name=?, metadata_value=? WHERE metadata_code = ?")

	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(metadata_name, metadata_value, metadata_code)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "[true,]")
}

// 现阶段直接移除，不做任何判断，以后要根据任务的状态以及分配情况选择直接删除还是标记删除
func RemoveMetadata(w http.ResponseWriter, r *http.Request) {
	metadata_code := strings.Join(r.Form["metadata_code"], "")

	fmt.Println("Remove mission metadata.", metadata_code)

	stmt, err := mydb.DBConn.Prepare("DELETE FROM mission_metadata WHERE metadata_code=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(metadata_code)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "[true,]")
}

func QueryMetadataByCode(metadata_code string) (metadata MissionMetadata) {
	stmt, err := mydb.DBConn.Prepare("SELECT metadata_code, mission_code, metadata_name, metadata_value FROM mission_metadata WHERE metadata_code=?")
	defer stmt.Close()
	if err != nil {
		panic(err.Error())
	}

	result, err := stmt.Query(metadata_code)
	defer result.Close()
	if err != nil {
		panic(err.Error())
	}

	//var project Project
	if result.Next() {
		err = result.Scan(&(metadata.MetadataCode), &(metadata.MissionCode), &(metadata.MetadataName), &(metadata.MetadataValue))
		if err != nil {
			panic(err.Error())
		}
	}
	return metadata
}

func GetMissionMetadata(w http.ResponseWriter, r *http.Request) {
	mission_code := strings.Join(r.Form["mission_code"], "")

	stmt, err := mydb.DBConn.Prepare("SELECT metadata_code, mission_code, metadata_name, metadata_value FROM mission_metadata WHERE mission_code=?")
	defer stmt.Close()
	if err != nil {
		panic(err.Error())
	}

	result, err := stmt.Query(mission_code)
	defer result.Close()
	if err != nil {
		panic(err.Error())
	}

	var metadata_list []interface{}

	for result.Next() {
		var metadata MissionMetadata
		err = result.Scan(&(metadata.MetadataCode), &(metadata.MissionCode), &(metadata.MetadataName), &(metadata.MetadataValue))
		if err != nil {
			panic(err.Error())
		}

		metadata_list = append(metadata_list, metadata)
	}

	res, _ := json.Marshal(metadata_list)

	fmt.Fprintf(w, string(res))
}
