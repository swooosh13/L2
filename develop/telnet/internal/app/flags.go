package app

import (
	"flag"
	"log"
	"strconv"
	"sync"
	"time"
)

type Flags struct {
	Timeout time.Duration
	Host    string
	Port    int
}

var flags *Flags
var once sync.Once

func GetFlags() *Flags {
	once.Do(func() {
		flags = &Flags{}
		flags.ParseFlags()
	})
	return flags
}

func (f *Flags) ParseFlags() {
	flag.DurationVar(&flags.Timeout, "timeout", 10 * time.Second, "set timeout connection")
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("error host type")
	}

	flags.Host = args[0]

	if len(args) > 1 {
		var err error
		flags.Port, err = strconv.Atoi(args[1])
		if err != nil {
			log.Fatal("error parsing port, invalid format")
		}
	} else {
		flags.Port = 23
	}

}
