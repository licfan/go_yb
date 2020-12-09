package utils

import (
	"github.com/go-playground/validator/v10"
)

func Verify(requset interface{}) interface{} {
	validate := validator.New()
	//validate.RegisterValidation("test", test)//自定义验证方法
	err := validate.Struct(requset)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return err
		}
	}
	return nil
}
