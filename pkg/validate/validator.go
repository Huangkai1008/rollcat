package validate

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
	"rollcat/pkg/e"
)

func init() {
	binding.Validator = new(defaultValidator)
}

type Register struct {
	Username string `json:"username" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email,max=256"`
}

type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (r *Register) Validate(errs validator.ValidationErrors) e.MarketError {
	var marketError e.MarketError
	err := errs[0] // 获得第一个错误并返回, yield错误

	if err.Field() == "Username" {
		switch err.Tag() {
		case "required":
			marketError.Message = "用户名不能为空"
		case "max":
			marketError.Message = fmt.Sprintf("用户名长度不能超过%s个字符", err.Param())
		}

	} else if err.Field() == "Password" {
		switch err.Tag() {
		case "required":
			marketError.Message = "密码不能为空"
		case "max":
			marketError.Message = fmt.Sprintf("密码长度不能超过%s个字符", err.Param())
		}
	} else if err.Field() == "Email" {
		switch err.Tag() {
		case "required":
			marketError.Message = "邮箱不能为空"
		case "email":
			marketError.Message = "邮箱格式不正确"
		case "max":
			marketError.Message = fmt.Sprintf("邮箱长度不能超过%s个字符", err.Param())
		}
	}

	return marketError
}

func (l *Login) Validate(errs validator.ValidationErrors) e.MarketError {
	var marketError e.MarketError
	err := errs[0] // 获得第一个错误并返回, yield错误
	if err.Field() == "Username" {
		switch err.Tag() {
		case "required":
			marketError.Message = "用户名不能为空"
		}

	} else if err.Field() == "Password" {
		switch err.Tag() {
		case "required":
			marketError.Message = "密码不能为空"
		}
	}

	return marketError

}
