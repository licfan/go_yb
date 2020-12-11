package utils

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"go.uber.org/zap"
	"go_yb/global"
)

func Verify(requset interface{}) string {
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()
	err := zhTranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		global.YB_LOG.Error("validate翻译", zap.Any("err", err))
	}
	//validate.RegisterValidation("test", test)//自定义验证方法
	err = validate.Struct(requset)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			//global.YB_LOG.Info(err.Translate(trans))
			return err.Translate(trans)
		}
	}
	return ""
}
