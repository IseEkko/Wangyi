package model

import "gorm.io/gorm"

//音乐详情
type Music_list struct {
	gorm.Model
	Name string //音乐名字
	Star string //作者
	Url  string //图片
	Time int    //歌曲时长度
	Img  string //海报图片
}
