package service

import (
	"singo/model"
	"singo/serializer"
)

type Music_list struct {
	Id   uint   `json:"id" form:"id"   `
	Name string `json:"name" form:"name"   binding:"required" `  //音乐名字
	Star string `json:"star" form:"star"    binding:"required" ` //作者
	Url  string `json:"url" form:"url"  binding:"required" `     //图片
	Time int    `json:"time" form:"time"  binding:"required" `   //歌曲时长度
	Img  string `json:"img" form:"img"  binding:"required"`      //海报图片
}

//创建歌曲使用
func (m Music_list) Creat_Musics() serializer.Response {
	music := model.Music_list{
		Name: m.Name,
		Star: m.Star,
		Url:  m.Url,
		Time: m.Time,
		Img:  m.Img,
	}
	res := model.DB.Create(&music)
	if res.Error != nil {
		return serializer.Json_Fail(100, "创建歌曲失败", res.Error)
	}
	return serializer.Json_Success(200, "歌曲创建成功", nil)
}
