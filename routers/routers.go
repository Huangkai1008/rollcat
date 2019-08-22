package routers

import (
	"github.com/gin-gonic/gin"
	"rollcat/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	return r
}
