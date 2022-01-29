package tourists

import (
	"cscke/internal/code"
	"cscke/internal/response"
	"cscke/internal/service"
	"cscke/internal/transformer"
	"cscke/internal/validate"
	"cscke/pkg/policy/context"
	"github.com/gin-gonic/gin"
)

// Authorized 获取授权地址
func Authorized(c *gin.Context) {

	params, err := validate.AuthorizedValidator(c)

	if err != nil {
		response.Code(c, code.ClientErr)
		return
	}

	response.Data(c, map[string]interface{}{
		"authorizedUrl": context.GetAuthContext(params.Platform).AuthorizedUrl(),
	})
}

// SnsLogin 授权登录
func SnsLogin(c *gin.Context) {

	params, err := validate.SnsLoginValidator(c)

	if err != nil {
		response.Code(c, code.ServerBusy)
		return
	}

	user, err := service.GetUserServ().LogSnsLogin(params.Platform, params.Ticket, c.ClientIP())

	if err != nil {
		response.Msg(c, err.Error())
		return
	}

	response.Data(c, map[string]interface{}{
		"user":  transformer.Item(user, transformer.NewUserDetailTransformer())["data"],
		"token": service.GetUserServ().GenTokenByUser(user),
	})
}
