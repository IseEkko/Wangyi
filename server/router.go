package server

import (
	"os"
	"singo/api"
	"singo/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())

	// 路由

	v1 := r.Group("/api/v1")
	{
		// 用户注册
		v1.POST("user/register", api.UserRegister)
		// 用户登录
		v1.POST("user/login", api.UserLogin)
		//用户通过token获取用户的信息
		v1.GET("user/user_info", middleware.JWTAuth(), api.User_info_Find)
		//修改头像的路径
		v1.POST("user/update_picurl", middleware.JWTAuth(), api.Update_pic)
		//修改用户的密码
		v1.POST("user/update_password", middleware.JWTAuth(), api.Update_Password)

	}
	//设备相关操作
	return r
}
