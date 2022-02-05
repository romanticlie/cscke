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

func TestFull(t *testing.T) {

	engine := boot.Setup()

	w := httptest.NewRecorder()

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxNjQzOTY2MjIzOTQ1ODYyNDcifQ.XyaDAeJbkKs9LwEPs4qaTz2CWJW3TjNy2gn7q79OUe0"

	req, err := fun.NewGetRequest("/api/user/full", map[string]string{"modules": ""}, map[string]string{"Authorization": token})

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

	if _, ok := data["user"]; !ok {
		t.Fatal("返回结果缺少user信息")
	}

}
