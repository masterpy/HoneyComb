//使用了第三方接口
// https://github.com/go-sql-driver/mysql

package mysqlUtility

import (
	_t "Honeycomb/utility"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DBConn *sql.DB

func init() {
	DBConn = ConnectToDB()
}

// TODO: 读入配置文件将来希望放到main中统一管理
func ConnectToDB() *sql.DB {
	if DBConn != nil {
		return DBConn
	}

	conf := "./Honeycomb.conf"
	propertyMap := _t.ReadConfig(conf)

	var userName, password, host, port, database string
	userName = propertyMap["DBUserName"]
	password = propertyMap["DBPassword"]
	host = propertyMap["DBIP"]
	port = propertyMap["DBPort"]
	database = propertyMap["DBDatabase"]

	sqlstring := userName + ":" + password + "@tcp(" + host + ":" + port + ")/" + database

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
