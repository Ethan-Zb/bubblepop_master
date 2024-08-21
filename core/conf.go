package core

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"

	"bubblepop_master/config"
	"bubblepop_master/global"
)

func InitConf() {
	const ConfigFile = "settings.yaml"

	c := &config.Config{}

	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yaml config error: %s", err))
	}

	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config init Unmarshal: %v", err)
	}

	log.Println("config init success")

	global.Config = c

}
