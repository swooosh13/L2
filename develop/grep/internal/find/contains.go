package find

import "strings"

func Contains(line string, pattern string, ignoreCase bool, fixed bool, flag *bool) {
	if fixed {
		if ignoreCase {
			if strings.EqualFold(strings.ToLower(line), strings.ToLower(pattern)) {
				*flag = true
			}
		} else {
			if strings.EqualFold(line, pattern) {
				*flag = true
			}
		}
	} else {
		if ignoreCase {
			if strings.Contains(strings.ToLower(line), strings.ToLower(pattern)) {
				*flag = true
			}
		} else {
			if strings.Contains(line, pattern) {
				*flag = true
			}
		}
	}
}

func IsContains(line string, pattern string, ignoreCase bool, fixed bool) bool {
	flag := false
	Contains(line, pattern, ignoreCase, fixed, &flag)
	return flag
}