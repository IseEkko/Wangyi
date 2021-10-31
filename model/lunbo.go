package model

import "gorm.io/gorm"

//轮播图相关
type Lunbo struct {
	gorm.Model
	Music_list_id uint       //绑定的music——id
	Music_list    Music_list //音乐详细
}
