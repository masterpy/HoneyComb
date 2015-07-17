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

type MissionNote struct {
	NoteCode    string `json:"note_code"`
	MissionCode string `json:"mission_code"`
	UserCode    string `json:"user_code"`
	NoteType    int    `json:"note_type"`
	Content     string `json:"content"`
}

func AddNote(w http.ResponseWriter, r *http.Request) {
	mission_code := strings.Join(r.Form["mission_code"], "")
	user_code := strings.Join(r.Form["user_code"], "")
	type_str := strings.Join(r.Form["note_type"], "")
	note_type, _ := strconv.Atoi(type_str)
	content := strings.Join(r.Form["content"], "")

	note_code := ut.GenerateCode(content)

	fmt.Println("Add Note for mission ", mission_code)

	stmt, err := mydb.DBConn.Prepare("INSERT INTO mission_note(note_code, mission_code, user_code, note_type, content) VALUES(?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(note_code, mission_code, user_code, note_type, content)
	if err != nil {
		panic(err.Error())
	}

	note := QueryNoteByCode(note_code)
	res, _ := json.Marshal(note)

	fmt.Fprintf(w, string(res))
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	note_code := strings.Join(r.Form["note_code"], "")
	user_code := strings.Join(r.Form["user_code"], "")
	type_str := strings.Join(r.Form["note_type"], "")
	note_type, _ := strconv.Atoi(type_str)
	content := strings.Join(r.Form["content"], "")

	fmt.Println("Update mission note", note_code)

	stmt, err := mydb.DBConn.Prepare("UPDATE mission_note SET user_code=?, note_type=?, content = ? WHERE note_code = ?")

	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user_code, note_type, content, note_code)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "[true,]")
}

// 现阶段直接移除，不做任何判断，以后要根据任务的状态以及分配情况选择直接删除还是标记删除
func RemoveNote(w http.ResponseWriter, r *http.Request) {
	note_code := strings.Join(r.Form["note_code"], "")

	fmt.Println("Remove mission note.", note_code)

	stmt, err := mydb.DBConn.Prepare("DELETE FROM mission_note WHERE note_code=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(note_code)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "[true,]")
}

func QueryNoteByCode(note_code string) (note MissionNote) {
	stmt, err := mydb.DBConn.Prepare("SELECT note_code, mission_code, user_code, note_type, content FROM mission_note WHERE note_code=?")
	defer stmt.Close()
	if err != nil {
		panic(err.Error())
	}

	result, err := stmt.Query(note_code)
	defer result.Close()
	if err != nil {
		panic(err.Error())
	}

	//var project Project
	if result.Next() {
		err = result.Scan(&(note.NoteCode), &(note.MissionCode), &(note.UserCode), &(note.NoteType), &(note.Content))
		if err != nil {
			panic(err.Error())
		}
	}
	return note
}

func GetMissionNote(w http.ResponseWriter, r *http.Request) {
	mission_code := strings.Join(r.Form["mission_code"], "")

	stmt, err := mydb.DBConn.Prepare("SELECT note_code, mission_code, user_code, note_type, content FROM mission_note WHERE mission_code=?")
	defer stmt.Close()
	if err != nil {
		panic(err.Error())
	}

	result, err := stmt.Query(mission_code)
	defer result.Close()
	if err != nil {
		panic(err.Error())
	}

	var note_list []interface{}

	for result.Next() {
		var note MissionNote
		err = result.Scan(&(note.NoteCode), &(note.MissionCode), &(note.UserCode), &(note.NoteType), &(note.Content))
		if err != nil {
			panic(err.Error())
		}

		note_list = append(note_list, note)
	}

	res, _ := json.Marshal(note_list)

	fmt.Fprintf(w, string(res))
}
