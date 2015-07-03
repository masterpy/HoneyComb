package configer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// 定义 MySQL 的配置类型
type MySQLConfig struct {
	Server   string `json:"mysql_user"`
	User     string `json:"mysql_password"`
	Password string `json:"mysql_server"`
	Database string `json:"mysql_database"`
}

var MySqlConf MySQLConfig

func init() {
	file_data, err := ioutil.ReadFile("./Honeycomb.conf")
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	var f interface{}
	err = json.Unmarshal(file_data, &f)
	property_map := f.(map[string]interface{})

	MySqlConf.User = property_map["mysql_user"].(string)
	MySqlConf.Password = property_map["mysql_password"].(string)
	MySqlConf.Server = property_map["mysql_server"].(string)
	MySqlConf.Database = property_map["mysql_database"].(string)
}
