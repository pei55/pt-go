package main

import (
	"github.com/pei55/pt-go/commands"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	commands.Execute()
}
