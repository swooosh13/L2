package app

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	flags2 "sortutil/flags"
	"sortutil/lines"
	"strings"
)

type App struct {
	flags     *flags2.Flags
	readPath  string
	writePath string
	data      []string
	sorted    bool
}

func GetApp(flags *flags2.Flags, rFile, wFile string) *App {
	return &App{
		flags:     flags,
		readPath:  rFile,
		writePath: wFile,
	}
}

func (a *App) LoadData() error {
	f, err := os.Open(a.readPath)
	if err != nil {
		return fmt.Errorf("error open file <%s>, error: %s", a.readPath, err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return fmt.Errorf("error read file <%s>, error: %s", a.readPath, err)
	}

	a.data = strings.Split(string(data), "\n")
	// игнорировать хвостовые пробелы
	if !a.flags.B {
		l := len(a.data)
		for i := 0; i < l; i++ {
			a.data[i] = strings.Trim(a.data[i], " ")
		}
	}

	// проверка отсортированы ли данные
	if a.flags.C {
		a.sorted = sort.StringsAreSorted(a.data)
	}

	return nil
}

func (a *App) WriteData() error {
	f, err := os.Create(a.writePath)
	if err != nil {
		return fmt.Errorf("error writing to file <%s>, error: %s", a.writePath, err)
	}
	defer f.Close()

	for _, line := range a.data {
		if _, err := f.WriteString(line + "\n"); err != nil {
			return fmt.Errorf("error writing data, error: %s", err)
		}
	}

	return nil
}

func (a *App) SortData() {
	if a.flags.K > 0 {
		a.data = lines.SortByColumn(a.data, a.flags)
	} else if a.flags.N {
		a.data = lines.SortByN(a.data, a.flags)
	} else if a.flags.R {
		lines.ReverseSort(&a.data)
	} else {
		lines.Sort(&a.data)
	}

	if a.flags.U {
		a.data = lines.UniqLines(a.data)
	}
}

func (a *App) Print() {
	fmt.Println("\tSorted")
	fmt.Println("-----")
	for _, v := range a.data {
		fmt.Println(v)
	}

	fmt.Println("-----")
}
