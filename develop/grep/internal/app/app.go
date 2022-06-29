package app

import (
	"fmt"
	"grep/internal/find"
	"grep/internal/flags"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type App struct {
	flags    *flags.Flags
	readPath string
	writePath string
	pattern  string
	data     []string
}

func GetApp(flags *flags.Flags, readPath, writePath, pattern string) *App {
	return &App{
		flags:    flags,
		readPath: readPath,
		writePath: writePath,
		pattern:  pattern,
	}
}

func (a *App) LoadData() error {
	f, err := os.Open(a.readPath)
	if err != nil {
		return fmt.Errorf("error open file <%s>, error: %s", a.readPath, err)
	}
	defer f.Close()

	// Optional: change with scanner scan by line
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return fmt.Errorf("error read file <%s>, error: %s", a.readPath, err)
	}

	a.data = strings.Split(string(data), "\n")

	return nil
}

func (a *App) Run() {
	var res []string

	if a.flags.After > 0 {
		res, _ = find.After(a.data, a.pattern, a.flags)
	} else if a.flags.Before > 0 {
		res, _ = find.Before(a.data, a.pattern, a.flags)
	} else if a.flags.Context > 0 {
		res, _ = find.WithContext(a.data, a.pattern, a.flags)
	} else if a.flags.Count {
		res = append(res, strconv.Itoa(find.Count(a.data, a.pattern, a.flags)))
	} else if a.flags.VInvert {
		res = append(res, strconv.Itoa(find.InvertCount(a.data, a.pattern, a.flags)))
	} else {
		res, _ = find.All(a.data, a.pattern, a.flags)
	}

	a.data = res
}

func (a *App) WriteData() error {
	if a.writePath != "" {
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
	} else {
		var out strings.Builder
		out.WriteString("[" + a.readPath + "]\n")
		for _, line := range a.data {
			if _, err := out.WriteString(line + "\n"); err != nil {
				return fmt.Errorf("error writing data, error: %s", err)
			}
		}

		fmt.Fprintf(os.Stdout, "%v", out.String())
	}

	return nil
}

