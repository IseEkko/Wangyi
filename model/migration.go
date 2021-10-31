package model

//执行数据迁移

func migration() {
	// 自动迁移模式
	_ = DB.AutoMigrate(&User{})            //用户表
	_ = DB.AutoMigrate(&Music_list{})      //音乐库
	_ = DB.AutoMigrate(&Lunbo{})           //轮播图
	_ = DB.AutoMigrate(&User_music_love{}) //用户喜欢列表

}
