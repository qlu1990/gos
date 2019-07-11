package conf

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Mongodb struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Server struct {
	Address string `yaml:"address"`
}
type Conf struct {
	Mongodb Mongodb `yaml:"mongodb"`
	Server  Server  `yaml:"server"`
}

var Cfg Conf

func LoadConf() {
	workdir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	f, err := ioutil.ReadFile(filepath.Join(workdir, "conf.yml"))
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(f, &Cfg)
	if err != nil {
		panic(err)
	}
}
