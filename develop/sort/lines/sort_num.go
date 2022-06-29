package lines

import (
	"log"
	"sort"
	"sortutil/flags"
	"strconv"
)

func SortByN(lines []string, flags *flags.Flags) []string {
	rs := make(map[float64][]string, len(lines))
	keys := make([]float64, 0, len(lines))

	for _, line := range lines {
		row := GetRow(line, flags)
		word, err := strconv.ParseFloat("-inf", 32) // если колонки нет значение -infinity
		if err != nil {
			log.Fatal(err)
		}
		if flags.K <= len(row) {
			word, err = strconv.ParseFloat(row[flags.K-1], 32) // парсим строку в число
			if err != nil {
				word, err = strconv.ParseFloat("-inf", 32) // если не число, то приравниваем -inf для сортировки
				if err != nil {
					log.Fatal(err)
				}
			}
		}

		if _, ok := rs[word]; !ok {
			keys = append(keys, word)
		}

		rs[word] = append(rs[word], line)
	}

	if flags.R {
		sort.Sort(sort.Reverse(sort.Float64Slice(keys)))
	} else {
		sort.Float64s(keys)
	}

	result := GetResult[float64](keys, rs)

	return result
}
