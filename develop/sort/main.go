package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sortutil/app"
	"sortutil/flags"
)

//
//-k — указание колонки для сортировки +
//-n — сортировать по числовому значению +
//-r — сортировать в обратном порядке +
//-u — не выводить повторяющиеся строки +

//Дополнительное
//
//-M — сортировать по названию месяца -
//-b — игнорировать хвостовые пробелы + (Read, Sorts)
//-c — проверять отсортированы ли данные + (SortData)
//-h — сортировать по числовому значению с учётом суффиксов -

func main() {
	flags := flags.GetFlags()
	args := flag.Args()
	if l := len(args); l != 2 {
		fmt.Println("should be 2 file paths")
		os.Exit(1)
	}

	app := app.GetApp(flags, args[0], args[1])
	err := app.LoadData()
	if err != nil {
		log.Fatal(err)
	}
	app.SortData()
	app.WriteData()

	os.Exit(0)
}
