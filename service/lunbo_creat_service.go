package service

import (
	"singo/model"
	"singo/serializer"
)

type Lunbo struct {
	Id            uint
	Music_list_id uint        `json:"music_list_id"  ` //绑定的music——id
	Music_list    Music_lists //音乐详细
	Num           int         `json:"num" form:"num" `
}
type Music_lists struct {
	Id   uint   `json:"id" form:"id"   `
	Name string `json:"name" form:"name"  `   //音乐名字
	Star string `json:"star" form:"star"    ` //作者
	Url  string `json:"url" form:"url"   `    //图片
	Time int    `json:"time" form:"time" `    //歌曲时长度
	Img  string `json:"img" form:"img" `      //海报图片
}

//返回轮播图相关信息  无查询功能，直接展示全部,传入参数num，查询指定的条数
func (l *Lunbo) Find_lun_bo() serializer.Response {
	var lunbo []Lunbo
	var nums int
	if l.Num == 0 {
		nums = 4
	} else {
		nums = l.Num
	}
	res := model.DB.Limit(nums).Preload("Music_list").Find(&lunbo)
	if res.Error != nil {
		return serializer.Json_Fail(100, "轮播图查找失败", nil)
	}
	return serializer.Json_Success(200, "轮播图返回成功", lunbo)
}
