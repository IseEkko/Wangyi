package service

import (
	"singo/model"
	"singo/serializer"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	WorkNumber   string `form:"worknumber" json:"worknumber" binding:"required"`
	Password     string `form:"password" json:"password" binding:"required,min=8,max=40"`
	UserName     string `form:"username" json:"username" binding:"required"`
	User_type_id int    `form:"user_type_id" json:"user_type_id" binding:"required"`
	Email        string `form:"eamil" json:"eamil" binding:"required"`
	Phone        string `form:"phone" json:"phone" binding:"required"`
}
type Change_code struct {
	Id uint `form:"id" json:"id" binding:"required"`
}

// valid 验证表单
func (service *UserRegisterService) valid() *serializer.Response {

	count := int64(0)
	model.DB.Model(&model.User{}).Where("work_number = ?", service.WorkNumber).Count(&count)
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
		WorkNumber:   service.WorkNumber,
		User_type_id: service.User_type_id,
		UserName:     service.UserName,
		Email:        service.Email,
		Phone:        service.Phone,
		Code:         0,
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
	return serializer.Json_Success(200, "注册成功", serializer.BuildUser(user))
}

//改变账号状态
func (c *Change_code) Change_user_code() serializer.Response {
	user := &model.User{}
	user.ID = c.Id
	rsult := model.DB.Find(&user)
	if rsult.Error != nil {
		return serializer.ParamErr("查询用户失败", rsult.Error)
	}
	if user.Code == 0 {
		model.DB.Model(&user).Update("code", 1)
	} else {
		model.DB.Model(&user).Update("code", 0)
	}
	return serializer.Json_Success(200, "修改状态成功", nil)
}
