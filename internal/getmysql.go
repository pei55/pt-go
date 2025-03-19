/**
*@Author: pei5
*@Date: 2024/8/17 上午11:19
*@File: internal/getmysql.go
*@Version:
*@Description:
**/
package internal

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
	"xorm.io/xorm"
)

func ConnectMysql(opt *MysqlOptin) (*xorm.Engine, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", opt.User, opt.Pass, opt.Host, opt.Port, opt.Db)
	if opt.Timezone != "" {
		dsn = dsn + fmt.Sprintf("&loc=%s", url.QueryEscape(opt.Timezone))
	}
	//fmt.Println(dsn)
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return engine, nil
}
