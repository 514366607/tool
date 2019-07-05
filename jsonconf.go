package tools

import (
	"encoding/json"
	io "io/ioutil"
)

// JSONLoad 读取配置
func JSONLoad(filename string, v interface{}) {
	data, err := io.ReadFile(filename)
	if err != nil {
		return
	}

	datajson := []byte(data)
	err = json.Unmarshal(datajson, v)
	if err != nil {
		return
	}
}
