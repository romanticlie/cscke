package service

import (
	"cscke/internal/model"
	"cscke/internal/repository"
	"cscke/pkg/policy/context"
	"errors"
	"gorm.io/gorm"
	"strconv"
	"sync"
)

var (
	userServ     *UserService
	userServOnce sync.Once
)

// GetUserServ  获取用户service
func GetUserServ() *UserService {

	if userServ == nil {
		userServOnce.Do(func() {
			userServ = &UserService{}
		})
	}

	return userServ
}

type UserService struct {
}

// LogSnsLogin 授权登录
func (u *UserService) LogSnsLogin(plat string, ticket string, registerIp string) (user *model.User, err error) {

	//先获取第三方用户信息
	authUser, err := context.GetAuthContext(plat).UserFromTicket(ticket)

	if err != nil {
		return nil, err
	}

	platform := repository.GetUserPlatformRepo().MapDriverPlatform(plat)

	//查看用户是否存在
	userPlatform, err := repository.GetUserRepo().FindByOpenid(platform, authUser.Openid)

	//查找用户信息
	if err == nil {
		return repository.GetUserRepo().GetByUniqueId(userPlatform.Userid)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {

		//开启事务
		tx := repository.D.Begin()

		user, err = repository.GetUserRepo().Create(registerIp, authUser.Nickname, authUser.Avatar, "")

		if err != nil {
			tx.Rollback()
			return nil, err
		}

		//同步到user_platform
		err = repository.GetUserPlatformRepo().CreateUserPlatform(
			user.Userid,
			platform,
			authUser.Openid,
			authUser.UnionId,
		)

		if err != nil {
			tx.Rollback()
			return nil, err
		}

		tx.Commit()
	}

	return user, err
}

// GenTokenByUser 通过用户生成jwtToken
func (u *UserService) GenTokenByUser(user *model.User) string {

	if user == nil {
		return ""
	}

	token, err := GetJwtServ().GenerateToken(strconv.FormatUint(user.Userid, 10))

	if err != nil {
		return ""
	}

	return token
}

// TokenParseUser jwtToken 获取用户信息
func (u *UserService) TokenParseUser(jwtToken string) (*model.User, error) {

	sub, err := GetJwtServ().ParseToken(jwtToken)

	if err != nil {
		return nil, err
	}

	userid, err := strconv.ParseUint(sub, 10, 64)

	if err != nil {
		return nil, err
	}

	//根据userid 获取用户信息
	user, err := repository.GetUserRepo().GetByUniqueId(userid)

	if err != nil {
		return nil, err
	}

	return user, nil
}
