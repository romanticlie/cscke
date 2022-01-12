package test

import (
	"cscke/internal/migration"
	"cscke/internal/model"
	"testing"
)


func TestUser(t *testing.T){

	migration.UserUp()
}

func TestCreate(t *testing.T){


}

func TestConst(t *testing.T){

	t.Log(model.TypeWechat,model.TypeWeibo,model.TypeQQ,model.TypeApple)
}