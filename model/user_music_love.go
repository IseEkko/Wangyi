package model

import "gorm.io/gorm"

//用户喜欢的歌单
type User_music_love struct {
	gorm.Model
	User_id       uint
	Music_list_id uint
	User          User
	Music_list    Music_list
}
