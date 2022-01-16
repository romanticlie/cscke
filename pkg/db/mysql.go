package db

import (
	"cscke/pkg/fun"
	"cscke/pkg/logmsg"
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/url"
	"sync"
	"time"
)

var (
	sqlConn *gorm.DB
	sqlOnce sync.Once
	sqlErr  error
)

type MySqlCfg struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Charset  string `yaml:"charset"`
	Loc      string `yaml:"loc"`
}

// MysqlConnect 获取Mysql的连接
func MysqlConnect() (*gorm.DB, error) {

	if sqlConn == nil {
		sqlOnce.Do(func() {

			var v *viper.Viper

			v, sqlErr = fun.GetYamlCfg("mysql")

			if sqlErr != nil {
				log.Println(logmsg.MysqlReadErr, sqlErr)
				return
			}

			cfg := &MySqlCfg{}

			sqlErr = v.UnmarshalKey("default", cfg)

			if sqlErr != nil {
				log.Println(logmsg.MysqlReadErr, sqlErr)
				return
			}

			dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
				cfg.Username,
				cfg.Password,
				cfg.Host,
				cfg.Port,
				cfg.Database,
				cfg.Charset,
				url.QueryEscape(cfg.Loc),
			)

			sqlConn, sqlErr = gorm.Open(mysql.Open(dns), &gorm.Config{
				Logger: logger.Default.LogMode(logger.Warn),
			})

			if sqlErr != nil {
				log.Println(logmsg.MysqlConnErr, sqlErr)
				return
			}

			var sqbDB *sql.DB

			sqbDB, sqlErr = sqlConn.DB()

			if sqlErr != nil {
				log.Println(logmsg.MysqlConnErr, sqlErr)
				return
			}

			//设置连接池
			sqbDB.SetMaxIdleConns(10)
			sqbDB.SetMaxOpenConns(100)
			sqbDB.SetConnMaxLifetime(time.Hour)

		})
	}

	return sqlConn, sqlErr
}
