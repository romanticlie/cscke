package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cross(c *gin.Context) {

	c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Access-Control-Expose-Headers", "*")
	c.Header("Access-Control-Allow-Credentials", "true")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}

	c.Next()
}
