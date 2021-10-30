package service

import (
	"github.com/dgrijalva/jwt-go"
	"singo/middleware"
	"singo/model"
	"singo/serializer"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required"` //用户名
	Password string `form:"password" json:"password" binding:"required"`   //用户密码
}

// setSession 设置session
func (service *UserLoginService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}

// Login 用户登录函数,用户所有的信息都存入的token中，在中间件中可以取出
//使用全局的token返回机制，set中间件，在get中返回相应的字段
func (service *UserLoginService) Login(c *gin.Context) serializer.Response {
	var user model.User
	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return serializer.ParamErr("账号或密码错误", err)
	}
	if user.CheckPassword(service.Password) == false {
		return serializer.ParamErr("密码错误", nil)
	}
	j := middleware.NewJWT()
	//存入token的字段
	claims := model.CustomClaims{
		Id:       user.ID,
		UserName: user.UserName,
		Password: user.Password,
		Local:    user.Local,
		Sex:      user.Sex,
		Jie:      user.Jie,
		HeadUrl:  user.HeadUrl,
		Birth:    user.Birth,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),            //签名生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24, //30天过期
			Issuer:    "user-web",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		return serializer.ParamErr("生成token失败", nil)
	}

	data := serializer.LoginResponse{
		Data:  user,
		Token: token,
	}
	return serializer.Json_Success(200, "登录成功", data)
}
