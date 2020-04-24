package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

var APPConfig appCfg

//解析配置文件
func Init(path ...string) {
	pwd, _ := os.Getwd()
	var p string
	if len(path) == 0 {
		p = fmt.Sprintf("%s/%s", pwd, "config.yaml")
	} else {
		p = path[0]
	}
	if !utils.IsExist(p) {
		panic(fmt.Sprintf("%s does not exist.",p))
	}
	if b, err := utils.HandFile(p); err != nil {
		panic(fmt.Sprintf("%s loading error:%v",p,err))
	} else {
		if err = yaml.Unmarshal(b, &APPConfig); err != nil {
			panic(fmt.Sprintf("Unmarshal config error:%v",err))
		}
	}
}
