package user

import (
	"cscke/internal/response"
	"cscke/internal/transformer"
	"github.com/gin-gonic/gin"
	"strings"
)

// Full 获取用户信息
func Full(c *gin.Context) {

	var modules []string

	m := c.Query("modules")

	if m != "" {
		modules = strings.Split(m, ",")
	}

	response.Data(c, map[string]interface{}{
		"user": transformer.Item(
			c.MustGet("user"),
			transformer.NewUserDetailTransformer(modules...),
		)["data"],
	})
}
