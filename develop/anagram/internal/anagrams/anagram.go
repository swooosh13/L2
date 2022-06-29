package anagrams

import (
	"fmt"
	"sort"
	"strings"
)

type anagrams map[string]*[]string

func (a *anagrams) IsUniq(w string) bool {
	key := SortString(w)
	words, ok := (*a)[key]
	if !ok {
		return true
	}

	for _, word := range *words {
		if w == word {
			return false
		}
	}

	return true
}

func FindAnagrams(words *[]string) *map[string]*[]string {
	mp := make(anagrams, 0)

	for _, w := range *words {
		w := strings.ToLower(w)
		key := SortString(w)
		// если нет - то создаем
		if _, ok := mp[key]; !ok {
			mp[key] = &[]string{w}
		} else if mp.IsUniq(w) {
			// добавляем только уникальное
			*mp[key] = append(*mp[key], w)
		}
	}

	for k, words := range mp {
		// удаление множества анаграмм если множество состоит из 1-го эл-та
		if len(*words) == 1 {
			delete(mp, k)
			continue
		}

		// ключ - первое встретившееся значение из мн-ва (преобразовываем)
		tmp := words
		delete(mp, k)
		mp[(*words)[0]] = tmp
		// сортировка по возрастанию
		sort.Sort(sort.Reverse(sort.StringSlice(*tmp)))

	}

	return (*map[string]*[]string)(&mp)
}

func PrintAnagrams(mp *map[string]*[]string) {
	for _, v := range *mp {
		fmt.Println("[", (*v)[0], "]")
		for _, w := range *v {
			fmt.Print(w, " ")
		}
		fmt.Println()
	}
}
