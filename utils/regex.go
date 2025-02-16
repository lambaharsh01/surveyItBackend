package utils

import "regexp"

func RegexEmail(str string) bool {
	reg := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return reg.MatchString(str)
}
func RegexWeakPassword(str string) bool {
	reg := regexp.MustCompile(`^(.{0,7}|[A-Za-z]+|\d+|[@$!%*?&]+)$`)
	return reg.MatchString(str)
}
func RegexPhone(str string) bool {
	reg := regexp.MustCompile(`^[6-9]\d{9}$`)
	return reg.MatchString(str)
}
func RegexDate(str string) bool {
	reg := regexp.MustCompile(`^(19|20)[0-9]{2}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$`)
	return reg.MatchString(str)
}
