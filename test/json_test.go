package test

import (
	"encoding/json"
	"testing"
)

type Ret struct {
	Token string
}

func TestJson(t *testing.T) {

	r := new(Ret)

	j := `{"Token":"123","UnionId":"123"}`

	err := json.Unmarshal([]byte(j),r)

	t.Log(err)
	t.Log(r)
}
