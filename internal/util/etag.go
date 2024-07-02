package util

import (
	"strconv"
)

func EtagLiteral(s string) string {
	return "`" + strconv.Quote(s) + "`"
}
