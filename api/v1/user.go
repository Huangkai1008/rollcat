package v1

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"rollcat/models"
	"rollcat/pkg/utils"
	"rollcat/pkg/validate"
)

func Register(c *gin.Context) {
	/**
	用户注册
	username 用户名
	password 密码
	email 邮箱
	*/
	var register validate.Register
	if err := c.ShouldBindJSON(&register); err != nil {
		errs := err.(validator.ValidationErrors)
		c.AbortWithStatusJSON(http.StatusBadRequest, register.Validate(errs))
	} else {
		params := map[string]interface{}{
			"username": register.Username,
		}
		if exist := models.ExistUser(params); exist {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "存在相同的用户名",
			})
			return
		}
		params = map[string]interface{}{
			"email": register.Email,
		}
		if exist := models.ExistUser(params); exist {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "存在相同的邮箱账户",
			})
			return
		}

		user := models.User{
			Username:     register.Username,
			Email:        register.Email,
			HashPassword: utils.MD5(register.Password),
		}

		if user, err := models.CreateUser(user); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, user)
		}
	}
}

func GetToken(c *gin.Context) {
	/**
	已注册用户获取token
	*/
	var login validate.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		errs := err.(validator.ValidationErrors)
		c.AbortWithStatusJSON(http.StatusBadRequest, login.Validate(errs))
	} else {
		maps := map[string]interface{}{"username": login.Username}
		if user, err := models.QueryUser(maps); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		} else if user == (models.User{}) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "不存在的用户名",
			})
			return
		} else if utils.MD5(login.Password) != user.HashPassword {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "用户名和密码不匹配",
			})
			return
		} else {
			if token, err := utils.GenerateToken(user.ID, user.Username); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": "token生成错误",
				})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{
					"token": token,
				})
			}
		}
	}

}
