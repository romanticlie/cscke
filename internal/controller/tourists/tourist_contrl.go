package tourists

import (
	"cscke/internal/code"
	"cscke/internal/response"
	"cscke/internal/service"
	"cscke/internal/transformer"
	"cscke/pkg/policy/context"
	"github.com/gin-gonic/gin"
)

// Authorized 获取授权地址
func Authorized(c *gin.Context) {

	platform, ok := c.GetQuery("platform")

	if !ok {
		response.Code(c, code.ClientErr)
		return
	}

	response.Data(c, map[string]interface{}{
		"authorizedUrl": context.GetAuthContext(platform).AuthorizedUrl(),
	})
}

// SnsLogin 授权登录
func SnsLogin(c *gin.Context) {

	params := &struct {
		Platform string `json:"platform"`
		Ticket   string `json:"ticket"`
		State    string `json:"state"`
	}{}

	if err := c.BindJSON(params); err != nil {
		response.Code(c, code.ClientErr)
		return
	}

	if params.Platform == "" || params.Ticket == "" {
		response.Code(c, code.ClientErr)
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
