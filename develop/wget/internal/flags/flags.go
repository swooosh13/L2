package flags

import (
	"log"
	"os"
)

type Flags struct {
	RootPath string
	DirName string
	Recursive bool
	Level int
}

var flags *Flags

func init() {
	flags = &Flags{}
	var err error
	flags.RootPath, err = os.Getwd()
	if err != nil {
		log.Fatalln("error root path", err)
	}
}

func GetFlags() *Flags {
	return flags
}