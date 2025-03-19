/**
*@Author: pei5
*@Date: 2025/3/19 20:52
*@File: commands/root.go
*@Version:
*@Description:
**/
package commands

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "miaodongbaike",
	Short: "秒懂百科服务",
	Long:  "从生产系统获取任务推送到秒懂百科",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(runCmd)
	//scanBaikeCmd.Flags().Int64VarP(&endTime, "endTime", "e", 0, "结束扫描时间戳(倒序),默认: 0")
	//rootCmd.AddCommand(scanBaikeCmd)
	//rootCmd.AddCommand(testCmd)
	//rootCmd.AddCommand(autoCrawLemmaCmd)
	//scanLemmaCmd.Flags().Int64VarP(&endTime, "endTime", "e", 0, "结束扫描时间戳(倒序),默认: 0")
	//rootCmd.AddCommand(scanLemmaCmd)
	rootCmd.Version = "0.0.1"
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	//rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "V", false, "verbose output")
}
