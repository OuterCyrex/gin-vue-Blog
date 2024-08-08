package validator

import (
	"gin-vue-blog/utils/errmsg"
	"github.com/go-playground/locales/zh_Hans_CN"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"
	"github.com/sirupsen/logrus"
	"reflect"
)

func Validate(data interface{}) (string, int) {
	validate := validator.New()
	uni := ut.New(zh_Hans_CN.New())
	trans, _ := uni.GetTranslator("zh_Hans_CN")

	err := zh.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		logrus.Error(err)
	}
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		label := fld.Tag.Get("label")
		return label
	})
	err = validate.Struct(data)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), errmsg.ERROR
		}
	}
	return "", errmsg.SUCCESS
}
