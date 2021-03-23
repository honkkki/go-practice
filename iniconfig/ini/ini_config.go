package ini

import (
	"fmt"
	"strings"
)

type Config struct {
	*MysqlConfig  `ini:"mysql"`
	*ServerConfig `ini:"server"`
}

type MysqlConfig struct {
	Host     string `ini:"host"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	Port     int    `ini:"port"`
}

type ServerConfig struct {
	Ip   string `ini:"ip"`
	Port int    `ini:"port"`
}

func UnMarshal(data []byte, res interface{}) (err error) {
	configArr := strings.Split(string(data), "\n")
	configMap := make(map[string]string)
	for _, v := range configArr{
		if strings.Contains(v, "=") {
			v = strings.Replace(v, " ", "", -1)
			itemArr := strings.Split(v, "=")
			configMap[itemArr[0]] = itemArr[1]
		}
	}

	for m, _ := range configMap {
		fmt.Println(m)
	}

	return
}