package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

// -f работает только с -d
Поддержать флаги:
-f - "fields" - выбрать поля (колонки) // LIST type, перечень полей для вырезания
-d - "delimiter" - использовать другой разделитель // кастомный разделитель
-s - "separated" - только строки с разделителем // если нет разделителя тогда не выводить строку

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type intslice []int

func (i *intslice) String() string {
	return fmt.Sprintf("%d", *i)
}

func (i *intslice) Set(value string) error {
	tmp, err := strconv.Atoi(value)
	if err != nil {
		*i = append(*i, -1)
	} else {
		*i = append(*i, tmp)
	}
	return nil
}

type Flags struct {
	Fields    intslice
	Delimiter string
	Separated bool
}

func main() {
	f := Flags{}
	// -i 1 -i 4 . . . [1, 4]
	flag.Var(&f.Fields, "i", "List of integers")
	flag.StringVar(&f.Delimiter, "d", " ", "\"delimiter\" - использовать другой разделитель")
	flag.BoolVar(&f.Separated, "s", false, "\"separated\" - только строки с разделителем")
	flag.Parse()

	data := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		} else {
			words := SplitLine(scanner.Text(), &f)
			if f.Separated && len(words) == 1 {
				continue
			}

			out := FilterWords(words, &f)
			data = append(data, JoinStrings(out, &f))
		}
	}
	fmt.Fprint(os.Stdout, strings.Join(data, "\n"))
}

func FilterWords(words []string, f *Flags) []string {
	out := make([]string, 0)
	if l := len(f.Fields); l > 0 {
		for _, col := range f.Fields {
			if col < 0 {
				log.Fatal("Fields should me more than 0")
			}

			if len(words) < col {
				continue
			}

			out = append(out, words[col-1])
		}
	}

	return out
}

func SplitLine(line string, flags *Flags) []string {
	return strings.Split(line, flags.Delimiter)
}

func JoinStrings(ss []string, flags *Flags) string {
	return strings.Join(ss, flags.Delimiter)
}
