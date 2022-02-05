package validate

import (
	"errors"
	"github.com/gin-gonic/gin"
	"regexp"
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

// Telephone 手机号登录参数
type Telephone struct {
	Tel    string `json:"tel" binding:"required,len=11"`
	Random string `json:"random" binding:"required,len=6"`
}

// SnsLoginValidator 授权登录的参数验证
func SnsLoginValidator(c *gin.Context) (*SnsLogin, error) {

	params := new(SnsLogin)

	//先校验数据格式是否正确
	if err := c.ShouldBindJSON(params); err != nil {
		return nil, err
	}

	return params, nil
}

// TelephoneValidator 手机号登录参数验证
func TelephoneValidator(c *gin.Context) (*Telephone, error) {

	params := new(Telephone)

	//先校验数据格式是否正确
	if err := c.ShouldBindJSON(params); err != nil {
		return nil, err
	}

	pattern := "^1[3-9][0-9]{9}$"

	reg := regexp.MustCompile(pattern)

	if !reg.MatchString(params.Tel) {
		return nil, errors.New("手机号格式错误")
	}

	return params, nil
}
