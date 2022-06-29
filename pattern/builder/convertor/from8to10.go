package convertor

import "strconv"

type from8To10 struct {
	str string
}

func (c *from8To10) Convert(s string) {
	// проверка
	for _, r := range s {
		if r < '0' || r > '7' {
			c.str = "incorrect value"
			return
		}
	}

	power := 1
	number := 0

	l := len(s)
	for i := l - 1; i >= 0; i-- {
		number += int(s[i] - '0')*power
		power *= 8
	}

	c.str = ""
	for number > 0 {
		c.str = strconv.Itoa(number%10) + c.str
		number /= 10
	}
}

func (c *from8To10) GetResult() string {
	return c.str
}
