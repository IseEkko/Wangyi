package service

import (
	"github.com/gin-gonic/gin"
	"singo/model"
	"singo/serializer"
)

type UserId struct {
	ID uint `form:"id" json:"id" `
}
type Username struct {
	Name string `form:"name" json:"name"  `
}

type user struct {
	ID       uint   `json:"id" form:"id"`
	UserName string `form:"user_name" json:"user_name" binding:"required" ` //用户名
	Password string `form:"password" json:"password" binding:"required"`    //用户密码
	Local    string `form:"local" json:"local" binding:"required" `         //用户地区
	Sex      string `form:"sex" json:"sex" binding:"required"`              //用户性别
	Jie      string ` form:"jie" json:"jie" binding:"required"`             //用户简介
	HeadUrl  string ` form:"head_url" json:"head_url" binding:"required"`   //用户头像路径
	Birth    string `  form:"birth" json:"birth" binding:"required"`        //用户生日
}

//返回用户详情
func (serve *UserId) User_info_Find(c *gin.Context) serializer.Response {
	v, _ := c.Get("user_name")
	data, err := model.GetUser_Username(v)
	if err != nil {
		return serializer.Json_Fail(100, "用户查找失败", err)
	}
	users := &user{
		ID:       data.ID,
		UserName: data.UserName,
		Password: data.Password,
		Local:    data.Local,
		Sex:      data.Sex,
		Jie:      data.Jie,
		HeadUrl:  data.HeadUrl,
		Birth:    data.Birth,
	}
	return serializer.Json_Success(200, "用户信息查找成功", users)
}
