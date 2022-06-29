package main

import (
	"flag"
	"fmt"
	app2 "grep/internal/app"
	"grep/internal/flags"
	"os"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения +
-B - "before" печатать +N строк до совпадения +
-C - "context" (A+B) печатать ±N строк вокруг совпадения +

-c - "count" (количество строк) +
-i - "ignore-case" (игнорировать регистр) +

-v - "invert" (вместо совпадения, исключать) +

-F - "fixed", точное совпадение со строкой, не паттерн +
-n - "line num", печатать номер строки +

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	flags := flags.GetFlags()
	args := flag.Args()

	l := len(args)
	if l < 2 {
		fmt.Println("should 2 more arguments - <path> <pattern> (<out>)")
		os.Exit(1)
	}

	var app *app2.App
	if l == 3 {
		app = app2.GetApp(flags, args[0], args[2], args[1])
	} else {
		app = app2.GetApp(flags, args[0], "", args[1])
	}

	app.LoadData()
	app.Run()
	app.WriteData()
}
