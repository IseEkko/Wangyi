package serializer

import "singo/model"

// User 用户序列化器
type User struct {
	ID           uint   `json:"id"`
	UserName     string `json:"user_name"`
	User_type_id int    `json:"User_type_id"`
	WorkNumber   string `json:"work_number"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:           user.ID,
		UserName:     user.UserName,
		User_type_id: user.User_type_id,
		WorkNumber:   user.WorkNumber,
		Email:        user.Email,
		Phone:        user.Phone,
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}
