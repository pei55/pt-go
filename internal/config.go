/**
*@Author: pei5
*@Date: 2024/1/18 5:13 PM
*@File: internal/config.yaml.go
*@Version:
*@Description:
**/
package internal

import (
	"fmt"
	"github.com/pei55/pgotools/pdirectory"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

const (
	AppName = "baikepush"
)

var (
	MyViper   *viper.Viper
	MyAppConf AppConf
)

type AppConf struct {
	StopFile  string     `yaml:"stopfile" mapstructure:"stopfile"`
	LogInfo   LogInfo    `yaml:"log_info" mapstructure:"log_info"`
	RedisInfo RedisInfo  `yaml:"redis_info" mapstructure:"redis_info"`
	MysqlInfo MysqlOptin `json:"mysql_info" mapstructure:"mysql_info"`
}

type LogInfo struct {
	LogName      string `mapstructure:"log_name"`
	LogPath      string `mapstructure:"log_path"`
	LogLevel     string `mapstructure:"log_level"`
	LogMaxSize   int    `mapstructure:"log_max_size"`
	LogMaxBackup int    `mapstructure:"log_max_backup"`
	LogMaxAge    int    `mapstructure:"log_max_age"`
}

type RedisInfo struct {
	Host           string `json:"host"`
	Password       string `json:"password"`
	QuePrefix      string `json:"que_prefix" mapstructure:"que_prefix"`
	QueName        string `json:"que_name" mapstructure:"que_name"`
	BiaozhuQueName string `json:"biaozhu_que_name" mapstructure:"biaozhu_que_name"`
	Port           int    `json:"port"`
	Db             int    `json:"db"`
}

type MysqlOptin struct {
	User     string `json:"user"`
	Pass     string `json:"pass"`
	Host     string `json:"host"`
	Db       string `json:"db"`
	Timezone string `json:"timezone"`
	Port     int    `json:"port"`
	ShowSql  bool   `json:"show_sql" mapstructure:"show_sql"`
}

func init() {
	workDir, _ := pdirectory.WorkDir()
	MyViper = viper.New()
	MyViper.SetConfigName("config")
	MyViper.SetConfigType("yaml")
	MyViper.AddConfigPath(fmt.Sprintf("%s/conf", filepath.Dir(os.Args[0])))
	MyViper.AddConfigPath(fmt.Sprintf("%s/conf", workDir))
	MyViper.AddConfigPath(fmt.Sprintf("/etc/%s/", AppName))
	MyViper.AddConfigPath(fmt.Sprintf("$HOME/.%s", AppName))
	//本机测试用这个
	//MyViper.SetConfigName("config")
	//MyViper.SetConfigType("yaml")
	//服务器用这个
	//MyViper.SetConfigFile(fmt.Sprintf("%s/%s", workDir, "conf/config.yaml")) //如果用这个，参数是配置文件绝对路径
	//fmt.Printf("当前工作目录:%s,文件:%s\n", fmt.Sprintf("%s/conf", workDir), MyViper.ConfigFileUsed())
	err := MyViper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = MyViper.UnmarshalKey("app_conf", &MyAppConf)
	if err != nil {
		panic(err)
	}
}
