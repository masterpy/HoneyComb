//使用了第三方接口
// https://github.com/go-sql-driver/mysql

package mysqlUtility

import (
	"Honeycomb/configer"
	"database/sql"
	// "encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	// "io/ioutil"
)

var DBConn *sql.DB

func init() {
	DBConn = ConnectToDB()
}

// TODO: 读入配置文件将来希望放到main中统一管理
// 读入配置文件已经放在了 configer 模块中
func ConnectToDB() *sql.DB {
	if DBConn != nil {
		return DBConn
	}

	// file_data, err := ioutil.ReadFile("./Honeycomb.conf")
	// if err != nil {
	// 	fmt.Printf("error: %v", err)
	// }

	// var f interface{}
	// err = json.Unmarshal(file_data, &f)
	// property_map := f.(map[string]interface{})

	// user := property_map["mysql_user"].(string)
	// password := property_map["mysql_password"].(string)
	// server := property_map["mysql_server"].(string)
	// database := property_map["mysql_database"].(string)

	// sqlstring := user + ":" + password + "@tcp(" + server + ")/" + database

	sqlstring := configer.MySqlConf.User + ":" + configer.MySqlConf.Password + "@tcp(" + configer.MySqlConf.Server + ")/" + configer.MySqlConf.Database

	fmt.Println("Connecting... ", sqlstring)

	DBConn, err := sql.Open("mysql", sqlstring)
	if err != nil {
		fmt.Println("Connect mysql false.")
	}

	return DBConn
}

func CloseDBConnection() {
	if DBConn != nil {
		DBConn.Close()
		DBConn = nil
	}
}
