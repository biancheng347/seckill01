package base

import (
	"github.com/astaxie/beego/logs"
	"encoding/json"
	"fmt"
)

type LogsConf struct{
	LogPath  string
	LogLevel string
}

func (p *LogsConf)InitLogConfig() (err error) {
	if err = appConfigStringValue(&p.LogPath, "log_path"); err != nil {
		return
	}
	if err = appConfigStringValue(&p.LogLevel, "log_level"); err != nil {
		return
	}
	return
}

func (p LogsConf)InitLogger() (err error) {
	config := make(map[string]interface{})
	config["filename"] = p.LogPath
	config["level"] = convertLogLevel(p.LogLevel)

	configByte, err := json.Marshal(config)
	if err != nil {
		err = fmt.Errorf("initLogger configByte failed,err:%v", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configByte))
	return
}

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