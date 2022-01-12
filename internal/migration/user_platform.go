package migration

import (
	"fmt"
	"cscke/internal/model"
	"cscke/pkg/db"
)

func UserPlatformUp(){
	fmt.Println("migration user_platform start")

	sqlConn,err := db.MysqlConnect()

	if err != nil {
		fmt.Println("migration user_platform failed",err)
		return
	}

	if err = sqlConn.AutoMigrate(&model.UserPlatform{}); err != nil {
		fmt.Println("migration user_platform failed",err)
		return
	}

	fmt.Println("migration user_platform end")
}
