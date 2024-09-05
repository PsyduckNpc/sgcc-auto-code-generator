package main

import (
	"github.com/zeromicro/go-zero/core/load"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/tools/goctl/cmd"
)

func main() { //开始的入口
	logx.Disable()
	load.Disable()
	cmd.Execute()
}
