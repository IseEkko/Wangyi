package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"singo/serializer"
	"singo/service"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var service service.UserRegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		HandleValidatorError(c, err)
		return
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		HandleValidatorError(c, err)
	}
}

// UserLogin 用户登录接口
func User_info_Find(c *gin.Context) {
	var service service.UserId
	if err := c.ShouldBind(&service); err == nil {
		res := service.User_info_Find(c)
		c.JSON(200, res)
	} else {
		HandleValidatorError(c, err)
	}
}

//// UserMe 用户详情
//func UserMe(c *gin.Context) {
//	res := serializer.BuildUserResponse(*user)
//	c.JSON(200, res)
//}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}

//修改头像路径
func Update_pic(c *gin.Context) {
	var service service.Pic_Update
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c)
		c.JSON(200, res)
	} else {
		HandleValidatorError(c, err)
	}
}

//func User_Id_Find(c *gin.Context) {
//	var service service.UserId
//	if err := c.ShouldBind(&service); err == nil {
//		res := service.User_Id_Find(c)
//		c.JSON(200, res)
//	} else {
//		HandleValidatorError(c, err)
//	}
//}
//更改用户的密码
func Update_Password(c *gin.Context) {
	var service service.Password_Update
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update_Password(c)
		c.JSON(200, res)
	} else {
		HandleValidatorError(c, err)
	}
}

//修改用户的基本信息
func Update_userinfo(c *gin.Context) {
	var service service.Update_UserInfo
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update_userinfo(c)
		c.JSON(200, res)
	} else {
		HandleValidatorError(c, err)
	}
}
