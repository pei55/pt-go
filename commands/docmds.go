/**
*@Author: pei5
*@Date: 2025/3/19 20:55
*@File: commands/docmds.go
*@Version:
*@Description:
**/
package commands

import (
	"github.com/pei55/pgotools/pdirectory"
	"github.com/pei55/pt-go/internal"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"runtime/debug"
	"syscall"
)

func runDaemon(cmd *cobra.Command, args []string) {
	// 检查是否已经是守护进程
	if os.Getppid() != 1 {
		// 将命令行程序转换为守护进程
		cmdDaemon := exec.Command(os.Args[0], os.Args[1:]...)
		cmdDaemon.Stdin = nil
		cmdDaemon.Stdout = nil
		cmdDaemon.Stderr = nil
		// 编译windows程序需要去掉
		cmdDaemon.SysProcAttr = &syscall.SysProcAttr{Setsid: true}

		_ = cmdDaemon.Start()
		os.Exit(0)
	}
	defer func() {
		if r := recover(); r != nil {
			internal.Log.Fatalf("Recovered from panic: %v\nStack trace: %s", r, string(debug.Stack()))
		}
	}()
	// 在这里执行实际的守护进程工作
	if pdirectory.IsExist(internal.MyAppConf.StopFile) {
		_ = os.Remove(internal.MyAppConf.StopFile)
	}
	if internal.MyAppConf.RedisInfo.Host == "" {
		internal.Log.Errorf("没有读取到redis配置")
		panic("没有读取到redis配置")
	}
	//internal.Log.Debugf(fmt.Sprintf("载入配置:%+v", internal.MyAppConf))
	internal.Log.Debugf("后台服务启动...")
	internal.Log.Debugf("后台进程id:%d", os.Getpid())

	// 这里启动后台程序用以标记词条状态
	for {
		// 这里运行服务程序
		break
	}
	internal.Log.Debugf("程序退出")
}

func aa(cmd *cobra.Command, args []string) {

}
