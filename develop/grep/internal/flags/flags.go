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
	After int
	Before int
	Context int
	Count bool
	Ignore bool
	VInvert bool
	Fixed bool
	LineNum bool
}

func (f *Flags) ParseFlags() {
	flag.IntVar(&f.After, "A", 0, "\"after\" печатать +N строк после совпадения")
	flag.IntVar(&f.Before, "B", 0, "\"before\" печатать +N строк до совпадения")
	flag.IntVar(&f.Context, "C", 0, "\"context\" (A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&f.Count, "c", false, "\"count\" (количество строк)")
	flag.BoolVar(&f.Ignore, "i", false, "\"ignore-case\" (игнорировать регистр)")
	flag.BoolVar(&f.VInvert, "v", false, "\"invert\" (вместо совпадения, исключать)")
	flag.BoolVar(&f.Fixed, "F", false, "\"fixed\", точное совпадение со строкой, не паттерн")
	flag.BoolVar(&f.LineNum, "n", false, "\"line num\", печатать номер строки")
	flag.Parse()
}
