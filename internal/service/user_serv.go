package service

import (
	"log"
	"cscke/internal/model"
	"cscke/internal/repository"
	"cscke/pkg/logmsg"
	"strconv"
	"sync"
)

var (
	userServ *UserService
	userServOnce sync.Once
)

// GetUserServ  获取用户service
func GetUserServ() *UserService{

	if userServ == nil {
		userServOnce.Do(func() {
			userServ = &UserService{}
		})
	}

	return userServ
}

type UserService struct {

}

// TokenParseUser jwtToken 获取用户信息
func (u *UserService) TokenParseUser(jwtToken string) (*model.User,error){

	sub,err := GetJwtServ().ParseToken(jwtToken)

	if err != nil {
		return nil,err
	}

	userid,err := strconv.ParseUint(sub,10,64)

	if err != nil {
		return nil,err
	}

	//根据userid 获取用户信息
	user,err := repository.GetUserRepo().GetByUniqueId(userid)

	if err != nil {
		return nil,err
	}

	return user,nil
}

// LoginByNickname 根据昵称登录
func (u *UserService) LoginByNickname(nickname string) (string,error) {

	user,err := repository.GetUserRepo().FindByNickname(nickname)

	if err != nil {
		return "",err
	}

	//缓存到redis
	_,err  = repository.GetUserRedisRepo().CacheByUserid(user.Userid,user)

	if err != nil {
		log.Println(logmsg.UserCacheErr,err)
	}

	//生成jwtToken
	return GetJwtServ().GenerateToken(strconv.FormatUint(user.Userid,10))
}

// CreateUser 创建用户
func (u *UserService)CreateUser(nickname string,gender int) bool {



	return true
}


