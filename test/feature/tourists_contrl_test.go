package feature

import (
	"cscke/internal/code"
	"cscke/internal/response"
	"cscke/pkg/boot"
	"cscke/pkg/fun"
	"encoding/json"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestTelephone 测试手机号注册登录
func TestTelephone(t *testing.T) {

	engine := boot.Setup()

	w := httptest.NewRecorder()

	//req, err := http.NewRequest("POST", "/api/tourists/telephone", nil)
	params, _ := json.Marshal(map[string]string{"tel": "13983517354", "random": "123456"})

	req, err := fun.NewJsonRequest("/api/tourists/telephone", params, nil)

	if err != nil {
		t.Fatal(err)
	}

	engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	resp := new(response.Response)

	if err = json.Unmarshal(w.Body.Bytes(), resp); err != nil {
		t.Fatal(err)
	}

	if resp.Code != code.None {
		t.Fatal(resp.Msg)
	}

	data, ok := resp.Data.(map[string]interface{})

	if !ok {
		t.Fatal("data 数据格式错误")
	}

	if _, ok := data["token"]; !ok {
		t.Fatal("返回结果缺少token")
	}
}

// TestSnsLogin 测试授权登录
//func TestSnsLogin(t *testing.T){
//
//	engine := boot.Setup()
//
//	w := httptest.NewRecorder()
//
//	req, _ := http.NewRequest("POST", "/api/tourists/snsLog", nil)
//	engine.ServeHTTP(w, req)
//
//	assert.Equal(t,http.StatusOK,w.Code)
//	t.Log(w.Body.String())
//}
