package repository

import (
	"encoding/json"
	"log"
	"cscke/internal/model"
	"cscke/pkg/fun"
	"cscke/pkg/logmsg"
	"strconv"
	"sync"
	"time"
)

var (
	userRepo     *UserRepo
	userRepoOnce sync.Once
)

// GetUserRepo 获取userModel数据仓库
func GetUserRepo() *UserRepo {

	if userRepo == nil {
		userRepoOnce.Do(func() {
			userRepo = &UserRepo{}
		})
	}

	return userRepo
}

type UserRepo struct {
}

// GetByUniqueId 根据主键获取用户信息
func (u *UserRepo) GetByUniqueId(userid uint64) (*model.User, error) {

	user := &model.User{}

	//先从缓存中读取
	j, err := GetUserRedisRepo().GetByUserid(userid)

	if err == nil {
		if err = json.Unmarshal([]byte(j), user); err == nil {
			return user, nil
		}
	}

	tx := d.First(user, userid)

	if tx.Error != nil {
		return user, err
	}

	//缓存用户信息
	_, err = GetUserRedisRepo().CacheByUserid(user.Userid, user)

	if err != nil {
		log.Println(logmsg.UserCacheErr, err)
	}

	return user, nil
}

// FindByNickname 根据nickname 获取用户信息
func (u *UserRepo) FindByNickname(nickname string) (*model.User, error) {

	user := &model.User{}

	tx := d.Where("nickname", nickname).First(user)

	return user, tx.Error
}

// Create 新建用户
func (u *UserRepo) Create(registerIp string,nickname string,telephone string) (user *model.User, err error) {

	uid, err := u.generateUserid()

	if err != nil {
		return nil, err
	}

	fields := []string{"Userid","Nickname","Avatar","Gender"}

	user = &model.User{
		Userid:     uid,
		Nickname:   nickname,
	}


	if telephone != "" {
		user.Telephone = telephone
		fields = append(fields,"Telephone")
	}


	tx := d.Select(fields).Create(user)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return user, nil
}

// generateUserid 生成随机的用户id
func (u *UserRepo) generateUserid() (uint64, error) {
	//uint64 1844674407370955161  19位

	//时间戳，最高1641545079186329，长度16位
	tp := time.Now().UnixMicro()

	//随机的两位
	rdNum := fun.Random(10, 99)

	sp := strconv.FormatInt(tp, 10) + strconv.Itoa(rdNum)

	return strconv.ParseUint(sp, 10, 64)
}
