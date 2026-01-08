package utils

import (
	"strconv"
)

func ConvertUnicodeString(s string) string {
	r, err := strconv.Unquote(`"` + s + `"`)
	if err != nil {
		return s
	}

	return r
}
