package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"cscke/internal/controller/resource"
	"cscke/internal/controller/tourist"
	"cscke/internal/controller/user"
	"cscke/internal/middleware"
)


func Boot(r *gin.Engine){

	//基础路由
	mapBaseRoute(r)

	//api路由
	mapApiRoute(r)
}

func mapBaseRoute(r *gin.Engine){

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK,"cscke")
	})

}

func mapApiRoute(r *gin.Engine){

	api := r.Group("/api")
	{
		ts := api.Group("/tourist")
		{
			ts.POST("/login",tourist.Login)
			ts.POST("/register",tourist.Register)
		}

		u := api.Group("/user")
		u.Use(middleware.Auth)
		{
			u.GET("/full",user.Full)
		}

		res := api.Group("/resource")
		{
			res.POST("/upload",resource.Upload)
		}
	}

}
