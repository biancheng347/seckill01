package config

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func convertLogLevel(level string) int {
	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}
	return logs.LevelDebug
}

func initLogger() (err error)  {
	config := make(map[string] interface{})
	config["filename"] = secKillConf.LogPath
	config["level"] = convertLogLevel(secKillConf.LogLevel)

	configByte,err := json.Marshal(config)
	if err != nil {
		err = fmt.Errorf("initLogger configByte failed,err:%v",err)
		return
	}

	logs.SetLogger(logs.AdapterFile,string(configByte))
	return
}


func InitSecKill() (err error) {

	err = initLogger()
	if err != nil {
		return
	}

	//err = initEtcd()
	//if err != nil {
	//	return
	//}
	//
	//err = loadSecConfig()
	//if err != nil {
	//	return
	//}
	//
	//models.InitServer(secKillConf)
	//initSecProcutWatcher()
	return
}