package util

import "strings"

const (
	EmptyString string = ""
)

func IsStringEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

func StringToLikeQueryExpression(s string) string {
	return "%" + s + "%"
}
