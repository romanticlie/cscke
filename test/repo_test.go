package test

import (
	"cscke/internal/repository"
	"testing"
)

func TestRepo(t *testing.T) {

	userPlatform, err := repository.GetUserRepo().FindByOpenid(1, "e1ac")

	t.Log(userPlatform, err)

	user, err := repository.GetUserRepo().GetByUniqueId(11)

	t.Log(user, err)

}
