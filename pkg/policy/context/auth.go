package context

import (
	"cscke/pkg/policy/contract"
	"cscke/pkg/policy/strategy"
)

const (
	WechatWeb = "wechatweb" //微信网页
)

var driverOpts = map[string]interface{}{
	WechatWeb: new(strategy.WechatWebStrategy),
}

type AuthContext struct {
	driver contract.AuthContract
}

// GetAuthContext 获取策略的实例
func GetAuthContext(platform string) *AuthContext {

	a := &AuthContext{}

	if driver, ok := driverOpts[platform]; ok {
		a.driver = driver.(contract.AuthContract)
	}

	return a
}

// AuthorizedUrl 生成授权登录的链接地址
func (a *AuthContext) AuthorizedUrl(state ...string) string {

	st := "STATE"

	if len(state) > 0 {
		st = state[0]
	}

	return a.driver.AccessUrl(st)
}

// UserFromTicket 根据code或者token 获取用户信息
func (a *AuthContext) UserFromTicket(ticket string) (*contract.AuthUser, error) {

	return a.driver.UserFromTicket(ticket)
}

// GetUser 获取用户信息
func (a *AuthContext) GetUser() *contract.AuthUser {

	return a.driver.GetUser()
}
