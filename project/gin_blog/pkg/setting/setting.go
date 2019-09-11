package setting

import (
    "io/ioutil"
    "gopkg.in/yaml.v2"
)

type mysql struct {
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
    Host string `yaml:"host"`
    DB string `yaml:"db"`
}

type upload struct {
	ImageUploadPath string `yaml:"imageUploadPath"`
	ImageMaxSize string `yaml:"imageMaxSize"`
	ImageAllowExts []string `yaml:"imageAllowExts"`
}

//Config系统配置配置
type config struct {
    Upload upload `yaml:"upload"`
    Mysql mysql `yaml:"mysql"`
}

var Config config

func init() {
    yamlFile, err := ioutil.ReadFile("conf/app.yaml")
    if err != nil {
        panic(err)
    }
    yaml.Unmarshal(yamlFile, &Config)
}