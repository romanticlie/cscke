// API服务的入口文件
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jpillora/overseer"
	"github.com/jpillora/overseer/fetcher"
	"cscke/internal/route"
	"cscke/pkg/db"
	"time"
)

const (
	Port = "8080" //服务监听端口
	BuildPath = "./build/api" //编译后的二进制文件路径
	CheckFreq = 3  //检查二进制文件频率间隔，单位秒
)

func main() {

	overseer.Run(overseer.Config{
		Program: Program,
		Address: ":" + Port,
		Fetcher: &fetcher.File{
			Path: BuildPath,
			Interval: time.Second * CheckFreq,
		},
	})
}


func Program(state overseer.State){

	r := gin.Default()

	route.Boot(r)

	//init db
	db.Boot()

	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	r.RunListener(state.Listener)
}
