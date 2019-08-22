package routers

import (
	"github.com/gin-gonic/gin"
	v1 "rollcat/api/v1"
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

	userApi := r.Group("/users")
	{
		userApi.POST("register", v1.Register)
		userApi.POST("tokens", v1.GetToken)
	}

	return r
}
