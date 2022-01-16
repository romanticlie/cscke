package repository

import (
	"cscke/internal/code"
	"cscke/internal/model"
	"encoding/json"
	"errors"
	"strconv"
	"sync"
	"time"
)

var (
	userRedisRepo     *UserRedisRepo
	userRedisRepoOnce sync.Once
)

// GetUserRedisRepo 获取userRedis数据仓库
func GetUserRedisRepo() *UserRedisRepo {

	if userRedisRepo == nil {
		userRedisRepoOnce.Do(func() {
			userRedisRepo = &UserRedisRepo{}
		})
	}

	return userRedisRepo
}

type UserRedisRepo struct {
}

func (u *UserRedisRepo) GetByUserid(userid uint64) (ret string, err error) {
	//redis 连接

	key := code.CacheKey(code.UserInfo, strconv.FormatUint(userid, 10))

	ret, err = Rd.Get(Rd.Context(), key).Result()

	if ret == "" {
		err = errors.New("userid not exists")
	}

	return
}

func (u *UserRedisRepo) CacheByUserid(userid uint64, user *model.User) (string, error) {

	//redis 连接

	key := code.CacheKey(code.UserInfo, strconv.FormatUint(userid, 10))

	j, err := json.Marshal(user)

	if err != nil {
		return "", err
	}

	return Rd.Set(Rd.Context(), key, j, time.Second*86400).Result()
}
