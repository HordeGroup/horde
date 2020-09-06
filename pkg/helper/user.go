package helper

import "regexp"

var (
	userNameRegexp, _ = regexp.Compile("^[a-zA-Z0-9]{4,16}$")
	passwordRegexp, _ = regexp.Compile("^[a-zA-Z0-9]{4,16}$")
)

func CheckUserName(name string) bool {
	return userNameRegexp.MatchString(name)
}

func CheckUserPwd(pwd string) bool {
	return passwordRegexp.MatchString(pwd)
}
