package contract

type AuthUser struct {
	Openid   string
	UnionId  string
	Nickname string
	Gender   int
	Avatar   string
}

type AuthOpts struct {
	AppId       string
	AppSecret   string
	RedirectUri string
}

type Token struct {
	AccessToken string
	Openid      string
	UnionId     string
}

type AuthContract interface {

	// GetConfig 获取基础配置
	GetConfig() *AuthOpts

	// AccessUrl 接入起始地址
	AccessUrl(state string) string

	// UserFromTicket 通过票据获取用户信息
	UserFromTicket(ticket string) (*AuthUser, error)

	// GetUser 获取用户信息
	GetUser() *AuthUser
}
