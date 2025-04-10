package c3po

import "strings"

type ValidationError struct {
	Rule  *Rule
	Field string
}

func (v ValidationError) Error() string {
	str := strings.Replace(v.Rule.Message, "{field}", "\""+v.Field+"\"", -1)
	str = strings.Replace(str, "{value}", "\""+v.Rule.Value+"\"", -1)
	return str
}

func IsValidationError(err error) bool {
	_, ok := err.(ValidationError)
	return ok
}
