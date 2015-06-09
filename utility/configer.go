package utility

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadConfig(fileName string) map[string]string {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		fmt.Println(fileName, err)
		return nil
	}
	buff := bufio.NewReader(file)
	propertyMap := make(map[string]string)
	for {
		line, err := buff.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		line = strings.Trim(line, "\n")
		propertyPair := strings.Split(line, "=")
		propertyMap[propertyPair[0]] = propertyPair[1]
	}
	return propertyMap
}

func OperationJson(result interface{}) []byte {
	data, _ := json.Marshal(result)
	return data
}
