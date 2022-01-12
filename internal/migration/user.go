package migration

import (
	"fmt"
	"cscke/internal/model"
	"cscke/pkg/db"
)

func UserUp(){
	fmt.Println("migration user start")

	sqlConn,err := db.MysqlConnect()

	if err != nil {
		fmt.Println("migration user failed",err)
		return
	}

	if err = sqlConn.AutoMigrate(&model.User{}); err != nil {
		fmt.Println("migration user failed",err)
		return
	}

	fmt.Println("migration user end")
}

