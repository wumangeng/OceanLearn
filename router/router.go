package router

import (
	"OceanLearn/auth"
	"OceanLearn/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	apiGroup := r.Group("/api")
	{
		userGroup := apiGroup.Group("/user")
		{
			userGroup.POST("/register", controller.RegisterUser)
			userGroup.POST("/login", controller.LoginUser)
			userGroup.GET("/info", auth.Auth(), controller.UserInfo) //请求需要登录才能访问
		}
	}
	return r

}
