package lines

import "sortutil/flags"

func SortByColumn(lines []string, flags *flags.Flags) []string {
	rs := make(map[string][]string, len(lines))
	keys := make([]string, len(lines))

	for _, line := range lines {
		row := GetRow(line, flags)

		var colWord string
		if flags.K <= len(row) - 1 {
			colWord = row[flags.K - 1]
		}

		if _, ok := rs[colWord]; !ok {
			keys = append(keys, colWord)
		}

		rs[colWord] = append(rs[colWord], line)
	}

	if flags.R {
		ReverseSort(&keys)
	} else {
		Sort(&keys)
	}

	result := GetResult[string](keys, rs)

	return result
}
