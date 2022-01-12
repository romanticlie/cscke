package db

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"log"
	"cscke/pkg/fun"
	"cscke/pkg/logmsg"
	"sync"
)

var (
	rdConn *redis.Client
	rdOnce sync.Once
	rdErr    error
)

type RdCfg struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

// RedisConnect 获取redis的连接
func RedisConnect() (*redis.Client, error) {

	if rdConn == nil {
		rdOnce.Do(func() {
			//获取配置
			var v *viper.Viper

			v, rdErr = fun.GetYamlCfg("redis")

			if rdErr != nil {
				log.Println(logmsg.RedisReadErr, rdErr.Error())
				return
			}

			cfg := &RdCfg{}

			rdErr = v.UnmarshalKey("default", cfg)

			if rdErr != nil {
				log.Println(logmsg.RedisReadErr, rdErr.Error())
				return
			}

			rdConn = redis.NewClient(&redis.Options{
				Addr:         cfg.Host + ":" + cfg.Port,
				Password:     cfg.Password, // no password set
				DB:           cfg.Db,       // use default DB
				MinIdleConns: 1,
			})

			rdErr = rdConn.Ping(rdConn.Context()).Err()
		})
	}

	return rdConn, rdErr
}
