/**
*@Author: pei5
*@Date: 2024/1/3 3:17 PM
*@File: internal/logger.go
*@Version:
*@Description:
**/
package internal

import (
	"github.com/pei55/pgotools/plog"
)

var (
	Log *plog.LogHandle
)

func init() {
	option := plog.NewLogOption(MyAppConf.LogInfo.LogMaxSize, MyAppConf.LogInfo.LogMaxBackup, MyAppConf.LogInfo.LogMaxAge, false, true)
	Log = plog.GetLogger(MyAppConf.LogInfo.LogName, MyAppConf.LogInfo.LogPath, option)
	// conf为空默认创建保留7天日志的配置
	if MyAppConf.LogInfo.LogLevel == "debug" {
		Log.Level = plog.LevelDebug
	} else {
		Log.Level = plog.LevelInfo
	}
}
