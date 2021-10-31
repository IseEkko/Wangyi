package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

//删除和创建喜欢
func Creat_user_love_music(c *gin.Context) {
	var service service.User_music_love
	if err := c.ShouldBind(&service); err == nil {
		res := service.Creat_user_love_music(c)
		c.JSON(200, res)
	} else {
		HandleValidatorError(c, err)
	}
}
