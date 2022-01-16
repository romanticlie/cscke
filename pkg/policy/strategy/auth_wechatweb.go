package strategy

import (
	"cscke/pkg/fun"
	"cscke/pkg/policy/contract"
	"encoding/json"
	"errors"
	"net/url"
	"sync"
)

var WechatWebOpts *contract.AuthOpts
var WechatWebOptsOnce sync.Once

type WechatWebStrategy struct {
	user      *contract.AuthUser
	accessUrl string
}

var _ contract.AuthContract = &WechatWebStrategy{}

func (w *WechatWebStrategy) GetConfig() *contract.AuthOpts {

	if WechatWebOpts == nil {

		WechatWebOptsOnce.Do(func() {
			c := &contract.AuthOpts{}

			v, _ := fun.GetYamlCfg("policy")

			if err := v.UnmarshalKey("authorization.wechatWeb", c); err == nil {
				WechatWebOpts = c
			}
		})
	}

	return WechatWebOpts
}

func (w *WechatWebStrategy) getAuthUrl() string {

	return "https://open.weixin.qq.com/connect/qrconnect"
}

// buildAuthUrl 拼接授权地址
func (w *WechatWebStrategy) buildAuthUrl(state string) string {

	c := w.GetConfig()

	p := url.Values{}
	p.Set("appid", c.AppId)
	p.Set("redirect_uri", c.RedirectUri)
	p.Set("response_type", "code")
	p.Set("scope", "snsapi_login")
	p.Set("state", state)

	return w.getAuthUrl() + "?" + p.Encode() + "#wechat_redirect"
}

// AccessUrl 获取授权地址
func (w *WechatWebStrategy) AccessUrl(state string) string {

	if w.accessUrl != "" {
		return w.accessUrl
	}

	w.accessUrl = w.buildAuthUrl(state)

	return w.accessUrl
}

func (w *WechatWebStrategy) getTokenUrl() string {

	return "https://api.weixin.qq.com/sns/oauth2/access_token"
}

// GetAccessToken 通过code获取accessToken
func (w *WechatWebStrategy) getAccessToken(code string) (t *contract.Token, err error) {

	c := w.GetConfig()

	resp, err := fun.HttpGet(w.getTokenUrl(), map[string]string{
		"appid":      c.AppId,
		"secret":     c.AppSecret,
		"code":       code,
		"grant_type": "authorization_code",
	}, nil)

	if err != nil {
		return nil, err
	}

	respFmt := &struct {
		AccessToken  string `json:"access_token"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		Openid       string `json:"openid"`
		Scope        string `json:"scope"`
		UnionId      string `json:"unionid"`
	}{}

	if err = json.Unmarshal(resp, respFmt); err != nil {
		return nil, errors.New(string(resp))
	}

	return &contract.Token{
		AccessToken: respFmt.AccessToken,
		Openid:      respFmt.Openid,
		UnionId:     respFmt.UnionId,
	}, nil
}

// getUserUrl 获取用户信息基础地址
func (w *WechatWebStrategy) getUserUrl() string {

	return "https://api.weixin.qq.com/sns/userinfo"
}

func (w *WechatWebStrategy) UserFromTicket(ticket string) (*contract.AuthUser, error) {

	token, err := w.getAccessToken(ticket)

	if err != nil {
		return nil, err
	}

	resp, err := fun.HttpGet(w.getUserUrl(), map[string]string{
		"access_token": token.AccessToken,
		"openid":       token.Openid,
		"lang":         "zh_CN",
	}, nil)

	respFmt := &struct {
		Openid     string   `json:"openid"`
		Nickname   string   `json:"nickname"`
		Sex        int      `json:"sex"`
		Province   string   `json:"province"`
		City       string   `json:"city"`
		Country    string   `json:"country"`
		HeadImgUrl string   `json:"headimgurl"`
		Privilege  []string `json:"privilege"`
		UnionId    string   `json:"unionid"`
	}{}

	if err = json.Unmarshal(resp, respFmt); err != nil {
		return nil, errors.New(string(resp))
	}

	w.user = &contract.AuthUser{
		Openid:   respFmt.Openid,
		UnionId:  respFmt.UnionId,
		Nickname: respFmt.Nickname,
		Gender:   respFmt.Sex,
		Avatar:   respFmt.HeadImgUrl,
	}

	return w.user, nil
}

// GetUser 获取用户信息
func (w *WechatWebStrategy) GetUser() *contract.AuthUser {

	return w.user
}
