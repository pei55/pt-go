/**
*@Author: pei5
*@Date: 2025/3/19 20:52
*@File: commands/cmds.go
*@Version:
*@Description:
**/
package commands

import "github.com/spf13/cobra"

var runCmd = &cobra.Command{
	Use:   "runbaike",
	Short: "运行服务",
	Long:  "运行百度秒懂百科服务",
	Run:   runDaemon,
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "测试",
	Run:   aa,
}
