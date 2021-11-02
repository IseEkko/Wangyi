package service

import (
	"github.com/gin-gonic/gin"
	"singo/model"
	"singo/serializer"
)

type User_music_love struct {
	Id            uint
	User_id       interface{}
	Music_list_id uint `json:"music_list_id" form:"music_list_id"   `
	Music         Music_list
	User          user
}

//type Music_list struct {
//	Id   uint   `json:"id" form:"id"   `
//	Name string `json:"name" form:"name"    ` //音乐名字
//	Star string `json:"star" form:"star"    ` //作者
//	Url  string `json:"url" form:"url"   `    //图片
//	Time int    `json:"time" form:"time"  `   //歌曲时长度
//	Img  string `json:"img" form:"img" `      //海报图片
//}

//创建用户的喜欢  删除喜欢
func (l *User_music_love) Creat_user_love_music(c *gin.Context) serializer.Response {
	if l.Music_list_id == 0 {
		return serializer.Json_Fail(422, "music_list_id不能为空", nil)
	}
	v, _ := c.Get("id")
	count := int64(0)
	love := &model.User_music_love{
		User_id:       v.(uint),
		Music_list_id: l.Music_list_id,
	}
	model.DB.Model(&model.User_music_love{}).Where("user_id = ? AND music_list_id = ?", v, l.Music_list_id).Count(&count)
	if count > 0 {
		res := model.DB.Where("user_id = ? AND music_list_id = ?", v, l.Music_list_id).Unscoped().Delete(&love)
		if res.Error != nil {
			return serializer.Json_Fail(100, "用户删除喜欢失败", nil)
		}
		return serializer.Json_Success(200, "用户删除喜欢成功", nil)
	} else {

		res := model.DB.Create(&love)
		if res.Error != nil {
			return serializer.Json_Fail(100, "用户添加喜欢失败", nil)
		}
		return serializer.Json_Success(200, "用户喜欢成功", nil)
	}
}

//展示用户喜欢
func Show_Love_Music(c *gin.Context) serializer.Response {
	v, _ := c.Get("id")
	var User_music_love []model.User_music_love
	res := model.DB.Preload("Music_list").Preload("User").Where("user_id", v).Find(&User_music_love)
	if res.Error != nil {
		return serializer.Json_Fail(100, "用户展示喜欢失败", nil)
	}
	return serializer.Json_Fail(100, "用户展示喜欢成功", User_music_love)
}
