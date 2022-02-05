package service

import (
	"cscke/internal/model"
	"cscke/internal/repository"
	"cscke/pkg/fun"
	"cscke/pkg/logmsg"
	"cscke/pkg/policy/context"
	"cscke/pkg/policy/contract"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"log"
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

type UserService struct{}

// GetUserByUserId 通过用户id获取用户信息
func (u *UserService) GetUserByUserId(userid uint64) (*model.User, error) {

	user := new(model.User)

	//先从缓存中读取
	j, err := repository.GetUserRedisRepo().GetByUserid(userid)

	if err != nil {
		return nil, err
	}

	if j != "" {
		//缓存序列化，直接返回
		if err = json.Unmarshal([]byte(j), user); err != nil {
			return nil, err
		}

		//直接返回
		return user, nil
	}

	//数据库查询
	user, err = repository.GetUserRepo().GetByUserid(userid)

	if err != nil {
		return nil, err
	}

	//缓存用户信息
	_, err = repository.GetUserRedisRepo().CacheByUserid(user.Userid, user)

	if err != nil {
		log.Println(logmsg.UserCacheErr, err)
		return nil, err
	}

	return user, nil
}

// registerUser 用户注册
func (u *UserService) registerUser(registerIp string, nickname string, avatar string, telephone string, platform int, authUser *contract.AuthUser) (*model.User, error) {

	//开启事务
	tx := repository.D.Begin()

	user, err := repository.GetUserRepo().Create(registerIp, nickname, avatar, telephone)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if platform != 0 {
		//如果是授权登录，同步到user_platform
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
	}

	tx.Commit()

	return user, nil
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
	userPlatform, err := repository.GetUserPlatformRepo().FindByOpenid(platform, authUser.Openid)

	//用户存在，直接获取用户
	if err == nil {
		return u.GetUserByUserId(userPlatform.Userid)
	}

	//如果用户不存在，注册一个用户
	if errors.Is(err, gorm.ErrRecordNotFound) {

		//开启事务
		user, err = u.registerUser(registerIp, authUser.Nickname, authUser.Avatar, "", platform, authUser)

		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}

	return user, err
}

// TelephoneLogin 手机号登录
func (u *UserService) TelephoneLogin(tel string, random string, registerIp string) (*model.User, error) {

	//检查验证码是否正确
	if !GetVerCodeServ().Check(VerCodeReg, tel, random) {
		return nil, errors.New("验证码错误")
	}

	//先查询当前手机号是否已经存在
	user, err := repository.GetUserRepo().FindByTelephone(tel)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		//新建用户

		user, err = u.registerUser(registerIp, u.RandomNickname(), u.DefaultAvatar(), tel, 0, nil)

		if err != nil {
			return nil, err
		}

	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

// RandomNickname 随机昵称
func (u *UserService) RandomNickname() string {
	return "cscke_" + strconv.Itoa(fun.Random(1, 9999))
}

// DefaultAvatar 默认头像
func (u *UserService) DefaultAvatar() string {
	return ""
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
	user, err := repository.GetUserRepo().GetByUserid(userid)

	if err != nil {
		return nil, err
	}

	return user, nil
}
