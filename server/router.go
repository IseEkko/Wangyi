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
		//修改用户的基本信息
		v1.POST("user/update_info", middleware.JWTAuth(), api.Update_userinfo)

		//创建歌曲
		v1.POST("music/creat_music", api.Creat_Musics)

		//轮播图展示传入num，num不是必须传入的，传入的num意思是需要多少条数据，默认为4
		v1.GET("lunbo/show", api.Find_lun_bo)

		//用户喜欢状态的修改
		v1.GET("music_love/change", middleware.JWTAuth(), api.Creat_user_love_music)
		//展示用户的喜欢歌曲
		v1.GET("music_love/show_love", middleware.JWTAuth(), api.Show_Love_Music)
		//展示新歌
		v1.GET("music_new/show", api.Find_new_music)

	}
	//设备相关操作
	return r
}
