package strategy

import (
	"encoding/json"
	"errors"
	"net/url"
	"cscke/pkg/fun"
	"cscke/pkg/policy/contract"
	"sync"
)

var WechatWebOpts *contract.AuthOpts
var WechatWebOptsOnce sync.Once

type WechatWebStrategy struct {
	token *contract.Token
	user *contract.AuthUser
}

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

func (w *WechatWebStrategy) GetAuthUrl() string {

	return "https://open.weixin.qq.com/connect/qrconnect"
}

func (w *WechatWebStrategy) BuildAuthUrl(state string) string {

	c := w.GetConfig()

	p := url.Values{}
	p.Set("appid", c.AppId)
	p.Set("redirect_uri", c.RedirectUri)
	p.Set("response_type", "code")
	p.Set("scope", "snsapi_login")
	p.Set("state", state)

	return w.GetAuthUrl() + "?" + p.Encode() + "#wechat_redirect"
}

func (w *WechatWebStrategy) GetTokenUrl() string {

	return "https://api.weixin.qq.com/sns/oauth2/access_token"
}

func (w *WechatWebStrategy) GetAccessToken(code string) (t *contract.Token, err error) {

	c := w.GetConfig()

	resp, err := fun.HttpGet(w.GetTokenUrl(), map[string]string{
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

	w.SetToken(&contract.Token{
		AccessToken: respFmt.AccessToken,
		Openid:      respFmt.Openid,
		UnionId:     respFmt.UnionId,
	})

	return w.GetToken(), nil
}

func (w *WechatWebStrategy) SetToken(token *contract.Token) {
	w.token = token
}

func (w *WechatWebStrategy) GetToken() *contract.Token{
	return w.token
}

func (w *WechatWebStrategy) GetUserUrl() string {

	return "https://api.weixin.qq.com/sns/userinfo"
}

func (w *WechatWebStrategy) GetUserByCode(code string) (*contract.AuthUser, error) {

	t, err := w.GetAccessToken(code)

	if err != nil {
		return nil, err
	}

	resp, err := fun.HttpGet(w.GetUserUrl(), map[string]string{
		"access_token": t.AccessToken,
		"openid":       t.Openid,
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

	w.SetUser(&contract.AuthUser{
		Openid:   respFmt.Openid,
		UnionId:  respFmt.UnionId,
		Nickname: respFmt.Nickname,
		Gender:   respFmt.Sex,
		Avatar:   respFmt.HeadImgUrl,
	})

	return w.GetUser(), nil
}

func (w *WechatWebStrategy) GetUserByToken(token string) (*contract.AuthUser, error) {
	return nil,nil
}

func (w *WechatWebStrategy) SetUser(user *contract.AuthUser){
	w.user = user
}

func (w *WechatWebStrategy) GetUser() *contract.AuthUser{
	return w.user
}



