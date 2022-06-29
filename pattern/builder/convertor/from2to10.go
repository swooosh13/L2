package convertor

import (
	"strconv"
)

// структура содержащая ср-ва конвертирования из 2-ой в 10-ую
type from2To10 struct {
	str string // число в 10-ой (конвертированное число)
}

// s - строка элементы которой в двоичной системе
func (c *from2To10) Convert(s string) {

	// проверка
	for _, r := range s {
		if int(r-'0') != 0 && int(r-'0') != 1 {
			c.str = "incorrect value"
			return
		}
	}

	power := 1
	number := 0

	l := len(s)
	for i := l - 1; i >= 0; i-- {
		if s[i] == '1' {
			number += power
		}
		power *= 2
	}

	c.str = ""
	for number > 0 {
		c.str = strconv.Itoa(number%10) + c.str
		number /= 10
	}
}

func (c *from2To10) GetResult() string {
	return c.str
}
