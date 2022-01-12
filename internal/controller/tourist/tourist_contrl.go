package tourist

import (
	"github.com/gin-gonic/gin"
	"cscke/internal/code"
	"cscke/internal/response"
	"cscke/internal/service"
	"cscke/internal/validate"
)


// Login 账号密码登录
func Login(c *gin.Context){


	nickname := c.PostForm("nickname")

	if nickname == "" {
		response.Code(c,code.ClientErr)
		return
	}

	token,err := service.GetUserServ().LoginByNickname(nickname)

	if err != nil {
		response.Code(c,code.LogFailed)
		return
	}

	response.Data(c, map[string]interface{}{
		"token": token,
	})
}


// SnsLog 第三方登录
func SnsLog(c *gin.Context){


	response.OK(c)
}

// Register 测试用户注册
func Register(c *gin.Context){

	v := &validate.Register{}

	if err := c.ShouldBindJSON(v); err != nil {
		response.Msg(c,err.Error())
		return
	}

	//创建用户
	if ok := service.GetUserServ().CreateUser(v.Nickname,v.Gender); !ok {
		response.Code(c,code.ServerBusy)
		return
	}

	response.OK(c)
}