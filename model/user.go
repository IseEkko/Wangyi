package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	UserName string //用户名
	Password string //用户密码
	Local    string //用户地区
	Sex      string //用户性别
	Jie      string //用户简介
	HeadUrl  string //用户头像路径
	Birth    string //用户生日
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
)

// GetUser 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

//根据用户的name查询数据
func GetUser_Username(Username interface{}) (User, error) {
	var user User
	result := DB.Where("user_name = ?", Username).First(&user)
	return user, result.Error
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
