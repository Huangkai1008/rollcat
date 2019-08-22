package routers

import (
	"github.com/gin-gonic/gin"
	"rollcat/middleware"
	"rollcat/pkg/logging"
	"rollcat/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.GinZap(logging.GinLogger))
	gin.SetMode(setting.RunMode)

	return r
}
