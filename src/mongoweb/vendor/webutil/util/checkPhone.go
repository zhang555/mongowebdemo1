package util

import "regexp"

func CheckPhoneNumber(number string) bool {
	reg := regexp.MustCompile(regular2)
	return reg.MatchString(number)
	//return false

}

const (
	regular  = "^(13[0-9]|14[57]|15[0-35-9]|18[07-9])\\d{8}$"
	regular2 = "^(13[0-9]|14[579]|15[0-3,5-9]|16[6]|17[0135678]|18[0-9]|19[89])\\d{8}$"
)

