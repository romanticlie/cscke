package contract



type AuthUser struct {
	Openid string
	UnionId string
	Nickname string
	Gender int
	Avatar string
}

type AuthOpts struct {
	AppId string
	AppSecret string
	RedirectUri string
}

type Token struct {
	AccessToken string
	Openid string
	UnionId string
}


type AuthContract interface {

	// GetConfig 获取基础配置
	GetConfig() *AuthOpts


	// GetAccessToken 获取accessToken
	GetAccessToken(code string) (t *Token, err error)


	// GetUser 获取用户信息
	GetUser() *AuthUser

}