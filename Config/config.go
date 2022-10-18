package Config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var Cfg struct {
	FrontendAddr string `json:"frontend_addr"` //前端监听地址
	BackendAddr  string `json:"backend_addr"`  //后端监听地址
}

func init() {
	GetConfig() //初始化配置结构体
}

//获取配置
func GetConfig() {

	//读取配置文件
	f, err := os.Open("./Config/.cfg.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)

	//获取配置
	json.Unmarshal(buf, &Cfg)

}
