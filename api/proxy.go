package api

import (
	"shProxy/serializer"

	"github.com/gin-gonic/gin"
)

//MyPing 测试
func MyPing(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 200,
		Msg:  "hello",
	})
}
