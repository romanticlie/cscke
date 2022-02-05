package route

import (
	"cscke/internal/controller/resource"
	"cscke/internal/controller/tourists"
	"cscke/internal/controller/user"
	"cscke/internal/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Boot(r *gin.Engine) {

	r.Use(middleware.Cross)

	//基础路由
	mapBaseRoute(r)

	//api路由
	mapApiRoute(r)
}

func mapBaseRoute(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "cscke")
	})

}

func mapApiRoute(r *gin.Engine) {

	api := r.Group("/api")
	{
		ts := api.Group("/tourists")
		{
			ts.GET("/authorized", tourists.Authorized)
			ts.POST("/snsLog", tourists.SnsLogin)
			ts.POST("/telephone", tourists.Telephone)
		}

		u := api.Group("/user")
		u.Use(middleware.Auth)
		{
			u.GET("/full", user.Full)
		}

		res := api.Group("/resource")
		{
			res.POST("/upload", resource.Upload)
		}
	}

}
