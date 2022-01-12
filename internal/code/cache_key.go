package code

import "strings"

const (
	UserInfo = "userinfo"
)


func CacheKey(k ...string) string{

	return strings.Join(k,":")
}
