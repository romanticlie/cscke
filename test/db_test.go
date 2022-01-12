package test

import (
	"fmt"
	"cscke/internal/model"
	"cscke/internal/repository"
	"cscke/pkg/db"
	"testing"
	"time"
)

func TestRedisConnect(t *testing.T) {

	rdConn, _ := db.RedisConnect()

	for i := 0; i < 20; i++ {
		go func() {
			num, _ := rdConn.Exists(rdConn.Context(), "zhang").Result()

			fmt.Println("num=", num)
		}()
	}

	time.Sleep(time.Second * 180)
}



func TestMysqlConnect(t *testing.T) {

	user := &model.User{}

	sqlConn,err := db.MysqlConnect()

	if err != nil {
		t.Fatal(err)
	}

	for i := 0;i < 20000;i++ {
		go func() {
			tx := sqlConn.First(user,1640598535)

			if tx.Error != nil {
				t.Fatal(tx.Error)
			}
		}()
	}

	time.Sleep(time.Second * 180)

}

func TestCacheUserid(t *testing.T){

	t.Log(repository.GetUserRedisRepo().GetByUserid(16405985351))
}
