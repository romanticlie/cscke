package repository

import "sync"

var (
	userPlatformRepo *UserPlatformRepo
	userPlatformRepoOnce sync.Once
)


func GetUserPlatformRepo() *UserPlatformRepo{
	if userPlatformRepo == nil {
		userPlatformRepoOnce.Do(func() {
			userPlatformRepo = &UserPlatformRepo{}
		})
	}

	return userPlatformRepo
}

type UserPlatformRepo struct {

}









