package validate

import (
	"github.com/gin-gonic/gin"
)

// Authorized 获取授权地址
type Authorized struct {
	Platform string `form:"platform" json:"platform" binding:"required,oneof=wechatweb"`
}

// AuthorizedValidator 获取授权地址参数
func AuthorizedValidator(c *gin.Context) (*Authorized, error) {
	params := new(Authorized)

	//先校验数据格式是否正确
	if err := c.ShouldBindQuery(params); err != nil {
		return nil, err
	}

	return params, nil
}

// SnsLogin 授权登录参数
type SnsLogin struct {
	Platform string `json:"platform" binding:"required,oneof=wechatweb"`
	Ticket   string `json:"ticket" binding:"required"`
	State    string `json:"state"`
}

func SnsLoginValidator(c *gin.Context) (*SnsLogin, error) {

	params := new(SnsLogin)

	//先校验数据格式是否正确
	if err := c.ShouldBindJSON(params); err != nil {
		return nil, err
	}

	return params, nil
}
