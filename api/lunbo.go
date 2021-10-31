package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

//轮播图展示
func Find_lun_bo(c *gin.Context) {
	var service service.Lunbo
	if err := c.ShouldBind(&service); err == nil {
		res := service.Find_lun_bo()
		c.JSON(200, res)
	} else {
		HandleValidatorError(c, err)
	}
}
