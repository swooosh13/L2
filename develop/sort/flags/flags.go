package flags

import (
	"flag"
	"sync"
)

var flags *Flags
var once sync.Once

func GetFlags() *Flags {
	once.Do(func() {
		flags = &Flags{}
		flags.ParseFlags()
	})
	return flags
}

type Flags struct {
	K int
	N bool
	R bool
	U bool

	M bool
	B bool
	C bool
	H bool
}

func (f *Flags) ParseFlags() {
	flag.IntVar(&f.K, "k", 0, "указание колонки для сортировки")
	flag.BoolVar(&f.N, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&f.R, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&f.U, "u", false, "не выводить повторяющиеся строки")
	flag.BoolVar(&f.M, "M", false, "сортировать по названию месяца")
	flag.BoolVar(&f.B, "b", false, "игнорировать хвостовые пробелы")
	flag.BoolVar(&f.C, "c", false, "проверять отсортированы ли данные")
	flag.BoolVar(&f.H, "h", false, "сортировать по числовому значению с учётом суффиксов")
	flag.Parse()
}
