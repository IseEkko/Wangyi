package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

//展示新的歌曲
func Find_new_music(c *gin.Context) {
	var service service.Limt
	if err := c.ShouldBind(&service); err == nil {
		res := service.Find_new_music()
		c.JSON(200, res)
	} else {
		HandleValidatorError(c, err)
	}
}
