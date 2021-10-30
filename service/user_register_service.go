package service

import (
	"singo/model"
	"singo/serializer"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	ID       uint   `json:"id" form:"id"`
	UserName string `form:"user_name" json:"user_name" binding:"required" ` //用户名
	Password string `form:"password" json:"password" binding:"required"`    //用户密码
	Local    string `form:"local" json:"local" binding:"required" `         //用户地区
	Sex      string `form:"sex" json:"sex" binding:"required"`              //用户性别
	Jie      string ` form:"jie" json:"jie" binding:"required"`             //用户简介
	HeadUrl  string ` form:"head_url" json:"head_url" binding:"required"`   //用户头像路径
	Birth    string `  form:"birth" json:"birth" binding:"required"`        //用户生日
}

// valid 验证表单
func (service *UserRegisterService) valid() *serializer.Response {

	count := int64(0)
	model.DB.Model(&model.User{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 416,
			Msg:  "账号已经注册占用",
		}
	}

	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() serializer.Response {
	user := model.User{
		UserName: service.UserName,
		Local:    service.Local,
		Sex:      service.Sex,
		Jie:      service.Jie,
		HeadUrl:  service.HeadUrl,
		Birth:    service.Birth,
	}
	// 表单验证
	if err := service.valid(); err != nil {
		return *err
	}
	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Err(
			serializer.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}
	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.ParamErr("注册失败", err)
	}
	return serializer.Json_Success(200, "注册成功", user)
}
