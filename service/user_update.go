package service

import (
	"github.com/gin-gonic/gin"
	"singo/model"
	"singo/serializer"
)

type Update_UserInfo struct {
	Local string `form:"local" json:"local"  ` //用户地区
	Sex   string `form:"sex" json:"sex" `      //用户性别
	Jie   string ` form:"jie" json:"jie"`      //用户简介
	Birth string `  form:"birth" json:"birth"` //用户生日
}

//修改用户的头像信息
type Pic_Update struct {
	Head_url string `form:"head_url" json:"head_url"  bind:"required"`
}

//修改用户的密码
type Password_Update struct {
	Old_password   string `form:"old_password" json:"old_password"  bind:"required"`     //旧密码
	Password       string `form:"password" json:"password"  bind:"required"`             //新密码
	Password_check string `form:"password_check" json:"password_check"  bind:"required"` //第二次密码
}

//修改用户头像信息url
func (pic *Pic_Update) Update(c *gin.Context) serializer.Response {
	if pic.Head_url == "" {
		return serializer.Json_Fail(422, "head_url 不能为空", nil)
	}
	v, _ := c.Get("user_name")
	res := model.DB.Model(model.User{}).Where("user_name", v).Update("head_url", pic.Head_url)
	if res.Error != nil {
		return serializer.Json_Fail(100, "修改图片路径失败", nil)
	}
	return serializer.Json_Success(200, "修改图片路径成功", nil)
}

//修改用户的密码
func (pass *Password_Update) Update_Password(c *gin.Context) serializer.Response {
	var user model.User
	v, _ := c.Get("user_name")
	res := model.DB.Model(model.User{}).Where("user_name", v).First(&user)
	if res.Error != nil {
		return serializer.Json_Fail(100, "获取用户密码失败", nil)
	}
	if !user.CheckPassword(pass.Old_password) {
		return serializer.Json_Fail(100, "用户原来的密码输入错误", nil)
	}
	if pass.Password != pass.Password_check {
		return serializer.Json_Fail(100, "用户两次新密码输入错误", nil)
	}
	if err := user.SetPassword(pass.Password); err != nil {
		return serializer.Err(
			serializer.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}
	res = model.DB.Model(model.User{}).Where("user_name", v).Update("password", user.Password)
	if res.Error != nil {
		return serializer.Json_Fail(100, "用户修改密码失败", nil)
	}
	return serializer.Json_Success(200, "修改密码成功", nil)
}

//修改用户的基本信息
func (u *Update_UserInfo) Update_userinfo(c *gin.Context) serializer.Response {
	v, _ := c.Get("user_name")
	User := &model.User{
		Password: u.Birth,
		Local:    u.Local,
		Sex:      u.Sex,
		Jie:      u.Jie,
		Birth:    u.Birth,
	}
	res := model.DB.Model(model.User{}).Where("user_name", v).Updates(User)
	if res.Error != nil {
		return serializer.Json_Fail(100, "用户修改信息失败", nil)
	}
	return serializer.Json_Success(200, "用户修改信息成功", nil)
}
