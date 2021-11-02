package service

import (
	"singo/model"
	"singo/serializer"
	"time"
)

type Music_list_new struct {
	Id   uint   `json:"id" form:"id"`
	Name string `json:"name" form:"name"` //音乐名字
	Star string `json:"star" form:"star"` //作者
	Url  string `json:"url" form:"url"`   //图片
	Time int    `json:"time" form:"time"` //歌曲时长度
	Img  string `json:"img" form:"img"`   //海报图片
}

//展示的条数
type Limt struct {
	Limts int `json:"limts" form:"limts"`
}

//展示新歌 展示的七天内的 传入limits 是查询的条数  不传入就是展示全部
func (m *Limt) Find_new_music() serializer.Response {
	var music []Music_list_new
	currentTime := time.Now()
	oldTime := currentTime.AddDate(0, 0, -7).Format("2006-01-02 15:04:05")
	num := m.Limts
	res := model.DB.Model(&model.Music_list{}).Limit(num).Where("created_at > ?", oldTime).Find(&music)
	if res.Error != nil {
		return serializer.Json_Fail(100, "展示新歌失败", nil)
	}
	return serializer.Json_Success(200, "展示新歌成功", music)
}
