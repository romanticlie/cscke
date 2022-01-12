package test

import (
	"cscke/pkg/fun"
	"testing"
	"time"
)

func TestTime(t *testing.T){

	t.Fatal(time.Now().UnixMicro())

	t.Log(time.Now())

	//时间戳
	t.Log(time.Now().Unix())

	tomorrow := time.Now().Add(time.Second * 86400)

	//Y-m-d H:i:s
	t.Log(tomorrow.Format("2006-01-02 15:04:05"))



	t.Log(time.Unix(time.Now().Unix(),0).Format("2006-01-02 15:04:05"))

	now := time.Now().Unix()
	t.Log(now)
	date := time.Unix(now,0).Format(fun.DateTime)
	t.Log(date)

	t.Log(time.ParseInLocation(fun.DateTime,date,time.Local))
}
