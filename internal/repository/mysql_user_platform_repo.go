package repository

import (
	"cscke/internal/model"
	"errors"
	"sync"
)

var (
	userPlatformRepo     *UserPlatformRepo
	userPlatformRepoOnce sync.Once
)

func GetUserPlatformRepo() *UserPlatformRepo {
	if userPlatformRepo == nil {
		userPlatformRepoOnce.Do(func() {
			userPlatformRepo = &UserPlatformRepo{}
		})
	}

	return userPlatformRepo
}

type UserPlatformRepo struct {
}

// MapDriverPlatform 映射platform
func (p *UserPlatformRepo) MapDriverPlatform(driverPlatform string) int {

	var driverPlatforms = map[string]int{
		"wechatweb": model.PlatFormWechatWeb,
		"wechatgzh": model.PlatFormWechatGZH,
		"wechatapp": model.PlatFormWechatApp,
	}

	platform, ok := driverPlatforms[driverPlatform]

	if !ok {
		return 0
	}

	return platform
}

func (p *UserPlatformRepo) CreateUserPlatform(userid uint64, platform int, openid string, unionId string) error {

	if platform == 0 || openid == "" {
		return errors.New("lack of platform or openid")
	}

	fields := []string{"Userid", "Platform", "Openid"}

	userPlatform := &model.UserPlatform{
		Userid:   userid,
		Platform: platform,
		Openid:   openid,
	}

	if unionId != "" {
		fields = append(fields, "UnionId")
	}

	tx := D.Select(fields).Create(userPlatform)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
