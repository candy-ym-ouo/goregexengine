package regex

import "regexp"

func HasBackreference(pattern string) bool {
	return regexp.MustCompile(`\\[1-9]|\\k<|\(\?P=`).MatchString(pattern)
}
