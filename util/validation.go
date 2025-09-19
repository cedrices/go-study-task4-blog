package util

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

// 错误信息过滤
func FilterEnglish(err error) string {
	// 类型断言：err 是否是 validator.ValidationErrors
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		errors := strings.Builder{}
		for _, fieldErr := range validationErrs {
			// key: 字段名，value: 错误原因
			errors.WriteString(fmt.Sprintf("%s: %s\n", fieldErr.Field(), getErrorMessage(fieldErr)))
		}
		return errors.String()
	}
	return err.Error()
}
func getErrorMessage(fieldErr validator.FieldError) string {
	switch fieldErr.Tag() {
	case "required":
		return "此字段为必填项"
	case "min":
		return fmt.Sprintf("长度不能少于 %s 个字符", fieldErr.Param())
	case "email":
		return "请输入有效的邮箱地址"
	default:
		return fmt.Sprintf("无效值")
	}
}
