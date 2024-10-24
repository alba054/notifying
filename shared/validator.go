package shared

import "strings"

func IsEmptyString(v string) bool {
	return strings.Trim(v, " ") == ""
}
