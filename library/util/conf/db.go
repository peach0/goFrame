package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type DatabaseConf struct {
	Mysqlurl   string `json:"mysqlurl"`
	Mysqlport  string `json:"mysqlport"`
	Mysqluname string `json:"mysqluname"`
	Mysqlupwd  string `json:"mysqlupwd"`
	Mysqldb    string `json:"mysqldb"`
	Redisurl   string `json:"redisurl"`
	Redisport  string `json:"redisport"`
	Redisupwd  string `json:"redisupwd"`
}

//读取配置文件，并解析json , 有错误抛异常，中止，不允许带病上岗
func LoadDatabaseConf() *DatabaseConf {

	dcobj := DatabaseConf{}
	file, err := os.Open("../conf/db.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fd, err := ioutil.ReadAll(file)
	conf := string(fd)
	//json解析到结构体里面
	err = json.Unmarshal([]byte(conf), &dcobj)
	if err != nil {
		panic(err)
	}
	return &dcobj
}
