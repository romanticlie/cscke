package repository

import (
	"cscke/internal/model"
	"cscke/pkg/fun"
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

// GetByUserid 根据主键获取用户信息
func (u *UserRepo) GetByUserid(userid uint64) (*model.User, error) {

	user := &model.User{}

	tx := D.Take(user, userid)

	if tx.Error != nil {
		return user, tx.Error
	}

	return user, nil
}

// FindByTelephone 通过手机号获取用户
func (u *UserRepo) FindByTelephone(tel string) (*model.User, error) {

	user := new(model.User)

	tx := D.Where("telephone = ?", tel).Take(user)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return user, nil
}

// Create 新建用户
func (u *UserRepo) Create(registerIp string, nickname string, avatar string, telephone string) (user *model.User, err error) {

	uid, err := u.generateUserid()

	if err != nil {
		return nil, err
	}

	fields := []string{"Userid", "Nickname", "Avatar", "RegisterIp"}

	user = &model.User{
		Userid:     uid,
		Nickname:   nickname,
		Avatar:     avatar,
		RegisterIp: registerIp,
	}

	if telephone != "" {
		user.Telephone = telephone
		fields = append(fields, "Telephone")
	}

	tx := D.Select(fields).Create(user)

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
