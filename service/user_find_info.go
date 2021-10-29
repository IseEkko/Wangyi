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
	Name       string
	Work       model.Types
	WorkNumber string
	Email      string
	Phone      string
}

//返回用户
func (serve *UserId) User_Id_Find(c *gin.Context) serializer.Response {
	v, _ := c.Get("userId")
	data, err := model.GetUser(v)
	if err != nil {
		return serializer.Json_Fail(100, "用户查找失败", err)
	}
	types, errs := model.GetTypes(data.User_type_id)
	users := &user{
		Name:       data.UserName,
		Work:       types,
		WorkNumber: data.WorkNumber,
		Email:      data.Email,
		Phone:      data.Phone,
	}
	if errs != nil {
		return serializer.Json_Fail(100, "用户type查找失败", err)
	}
	return serializer.Json_Success(200, "用户信息查找成功", users)
}

//查询所有的用户，如果输入名字，可以做查询
func (u *Username) User_all_show() serializer.Response {
	var users []model.User
	res := model.DB.Where(&model.User{UserName: u.Name}).Find(&users)
	if res.Error != nil {
		return serializer.Json_Fail(100, "用户查找失败", res.Error)
	}
	return serializer.Json_Success(200, "用户信息查找成功", users)
}
