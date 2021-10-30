package serializer

import "singo/model"

// User 用户序列化器
type User struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"` //用户名
	Password string `json:"password"`  //用户密码
	Local    string `json:"local"`     //用户地区
	Sex      string `json:"sex"`       //用户性别
	Jie      string `json:"jie"`       //用户简介
	HeadUrl  string `json:"head_url"`  //用户头像路径
	Birth    string `json:"birth"`     //用户生日
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:       user.ID,
		UserName: user.UserName,
		Password: user.Password,
		Local:    user.Local,
		Sex:      user.Sex,
		Jie:      user.Jie,
		HeadUrl:  user.HeadUrl,
		Birth:    user.Birth,
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}
