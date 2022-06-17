package config

import (
	"encoding/json"
	"github.com/TarsCloud/TarsGo/tars"
	"io/ioutil"
	"runtime"
)

var (
	Cfg              = new(CommonConfig)
	ConfPath string = ".\\"
)

type CommonConfig struct {
	*CommonConf
}

func InitCommonConfig() {
	var path = ""
	var err error
	if runtime.GOOS == "darwin" {
		path = "./config/"
	} else if runtime.GOOS == "windows" {
		path = ConfPath
	} else {
		path =  tars.GetServerConfig().BasePath
	}

	data, err := ioutil.ReadFile(path + "/common.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, Cfg)
	if err != nil {
		panic(err)
	}
}

func InitDBConfig() {
	// config
	var path = ""
	var err error
	if runtime.GOOS == "darwin" {
		path = "./config/"
	} else if runtime.GOOS == "windows" {
		path = ConfPath
	} else {
		path =  tars.GetServerConfig().BasePath
	}

	dataDb, err := ioutil.ReadFile(path + "db.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(dataDb, Cfg)
	if err != nil {
		panic(err)
	}

	//log.Info("加载配置完成", zap.Any("conf", Cfg))
}

