package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	codes "cscke/internal/code"
)

// OK 直接返回成功信息
func OK(c *gin.Context){
	Result(c,codes.None,make(map[string]interface{}),"")
}

// Code 只返回错误码信息
func Code(c *gin.Context,code int){

	Result(c,code,make(map[string]interface{}),"")
}

// Data 成功返回数据
func Data(c *gin.Context,data map[string]interface{}){
	Result(c,codes.None,data,"")
}


// Msg 只返回提示信息
func Msg(c *gin.Context,msg string){
	Result(c,codes.Prompt,make(map[string]interface{}),msg)
}

// Result 返回完整的结构信息
func Result(c *gin.Context,code int,data map[string]interface{},msg string){


	if errMsg,ok := codes.ErrCodeMessage[code]; ok {
		msg = errMsg
	}

	c.JSON(http.StatusOK,gin.H{
		"code": code,
		"data": data,
		"msg": msg,
	})
}
