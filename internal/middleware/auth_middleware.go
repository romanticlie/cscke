package middleware

import (
	"cscke/internal/code"
	"cscke/internal/response"
	"cscke/internal/service"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {

	token := c.GetHeader("Authorization")

	if token == "" {
		response.Code(c, code.AuthFailed)
		c.Abort()
	}

	user, err := service.GetUserServ().TokenParseUser(token)

	if err != nil {
		response.Code(c, code.AuthFailed)
		c.Abort()
	}

	//注入用户信息
	c.Set("user", user)

	c.Next()
}
