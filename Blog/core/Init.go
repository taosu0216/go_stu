package core

import (
	"Blog/config"
	"Blog/global"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

func InitConfig() {
	const ConfigFile = "config.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalln("Mysql连接失败,error: ", err)
	}
	log.Println("Config Inited successfully...")
	global.Config = c
}
