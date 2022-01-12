package test

import (
	"cscke/pkg/policy"
	"testing"
)

func TestPolicy(t *testing.T){

	var strategy policy.AuthContract = &policy.WechatWebStrategy{}

	t.Log(strategy)


}
