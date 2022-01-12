// 数据迁移
package main

import "cscke/internal/migration"

func main() {

	migration.UserUp()
	migration.UserPlatformUp()
}
