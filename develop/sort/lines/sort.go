package lines

import (
	"sort"
	"sortutil/flags"
	"strings"
)

func ReverseSort(lines *[]string) {
	sort.Sort(sort.Reverse(sort.StringSlice(*lines)))
}

func Sort(lines *[]string) {
	sort.Strings(*lines)
}

func IsSorted(lines []string) bool {
	return sort.StringsAreSorted(lines)
}

// out only uniq lines
// optimal for sorted rows

func UniqLines(lines []string) []string {
	set := make(map[string]struct{})
	uniq := make([]int, 0)

	for i, v := range lines {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
			uniq = append(uniq, i)
		}
	}

	result := make([]string, len(uniq))
	for i, v := range uniq {
		result[i] = lines[v]
	}

	return result
}

func GetRow(line string, flags *flags.Flags) []string {
	var row []string
	if flags.B {
		row = strings.Fields(line)
	} else {
		row = strings.Split(line, " ")
	}

	return row
}

func GetResult[T comparable](keys []T, rs map[T][]string) []string {
	result := make([]string, 0)

	for _, k := range keys {
		result = append(result, rs[k]...)
	}
	return result
}
