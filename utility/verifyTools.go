package utility

import (
	"regexp"
)

func VerifyPhone(phoneNum string) bool {
	reg := `^(0|86|17951)?(13[0-9]|15[012356789]|18[0-9]|14[57]|17[0-9])[0-9]{8}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(phoneNum)
}

func VerifyEmail(email string) bool {
	//sweetsmelon@163.com
	reg := `^[A-Za-z0-9]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(email)
}
func VerifyPassword(password string) bool {
	//密码必须
	// 1、以字母开头
	//2、长度为6-16
	//3、可以是字母或数字
	reg := `^[a-zA-Z]{1}[a-zA-Z0-9]{5,15}$`
	//reg := `[A-Za-z][0-9]{6,16}`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(password)
}
