package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func unpack(s string) string {
	rs := append([]rune(s), 'x')
	l := len(rs)

	// invalid
	if l == 0 {
		return ""
	}

	prev := rs[0]
	// invalid
	if unicode.IsDigit(prev) {
		return ""
	}

	result := strings.Builder{}
	isEscape := prev == '\\'
	for i := 1; i < l; i++ {
		curr := rs[i]

		// если число
		if unicode.IsDigit(curr) && unicode.IsLetter(prev) {
			// если предыдущее - символ => добавляем к строке
			for cnt, _ := strconv.Atoi(string(curr)); cnt > 0; cnt-- {
				result.WriteRune(prev)
			}
		}

		if !unicode.IsDigit(curr) && unicode.IsLetter(prev) {
			result.WriteRune(prev)
		}

		if unicode.IsDigit(prev) && isEscape {
			if unicode.IsDigit(curr) {
				for cnt, _ := strconv.Atoi(string(curr)); cnt > 0; cnt-- {
					result.WriteRune(prev)
				}
			}

			if unicode.IsLetter(curr) || curr == '\\' {
				result.WriteRune(prev)
			}

			isEscape = false
		}

		if curr == '\\' {
			isEscape = true
		}

		if prev == '\\' && i > 1 && rs[i-2] == '\\' {

			if unicode.IsDigit(curr) {
				for cnt, _ := strconv.Atoi(string(curr)); cnt > 0; cnt-- {
					result.WriteRune(prev)
				}
			} else {
				result.WriteRune(prev)
			}

			isEscape = false
		}

		prev = curr
	}

	return result.String()
}

func main() {
	//ss := `qwe\4\5`
	//ss := `qwe\45`
	//ss := `qwe\\5`
	//ss := `qwe\\5`
	ss := `\\4\\4\\`
	fmt.Println(unpack(ss))
}
