/*
Строитель используется для конвертации числа из
одной системы счисления в другую (2 -> 10, 8 -> 10)

Применимость
* когда хотим избавиться от огромного конструктора с большим числом параметров (или разбив его на 10 конструкторов)

+ Позволяет создавать продукты пошагово
+ Один и тот же код для создания различных продуктов
+ изоляция кода сборки продукта от его бизнес логики

- усложняет код из-за доп. классов
- клиент привязан к конкретным классам строителей, которых может не быть
 	метода получения результата в интерфейсе директора

*/

package convertor

import "errors"

type Builder interface {
	Convert(string)
	GetResult() string // string - продукт builder`а
}

func GetBuilder(from, to int) (Builder, error) {
	switch from {
	case 2:
		switch to {
		case 10:
			return &from2To10{}, nil
		}
	case 8:
		switch to {
		case 10:
			return &from8To10{}, nil
		}
	default:
		return nil, errors.New("builder does not exist")
	}

	return nil, errors.New("builder does not exist")
}