// API服务的入口文件
package main

import (
	"cscke/pkg/boot"
	"github.com/jpillora/overseer"
	"github.com/jpillora/overseer/fetcher"
	"time"
)

const (
	Port      = "8080"        //服务监听端口
	BuildPath = "./build/api" //编译后的二进制文件路径
	CheckFreq = 3             //检查二进制文件频率间隔，单位秒
)

func main() {

	overseer.Run(overseer.Config{
		Program: Program,
		Address: ":" + Port,
		Fetcher: &fetcher.File{
			Path:     BuildPath,
			Interval: time.Second * CheckFreq,
		},
	})
}

func Program(state overseer.State) {

	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	boot.Setup().RunListener(state.Listener)
}
