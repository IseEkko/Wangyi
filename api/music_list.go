package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

//创建音乐信息
func Creat_Musics(c *gin.Context) {
	var service service.Music_list
	if err := c.ShouldBind(&service); err == nil {
		res := service.Creat_Musics()
		c.JSON(200, res)
	} else {
		HandleValidatorError(c, err)
	}
}
