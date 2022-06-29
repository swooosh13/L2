package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	flags2 "wget/internal/flags"
	"wget/internal/utils"
)

func main() {
	flags := flags2.GetFlags()
	flag.StringVar(&flags.RootPath, "p", flags.RootPath, "set root path to download")
	flag.StringVar(&flags.DirName, "d", "sites", "set download directory name")
	flag.BoolVar(&flags.Recursive, "r", false, "download recursive")
	flag.IntVar(&flags.Level, "l", 0, "download recursive level")
	flag.Parse()
	args := flag.Args()
	// TODO remove
	//args = []string{"https://www.aon.com/home/index"}

	client := &http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	urls, err := utils.StringsToUrls(args)
	if err != nil {
		log.Fatal("invalid url:", err)
	}

	// concurrent with WG
	wg := &sync.WaitGroup{}
	for _, url := range urls {
		wg.Add(1)
		go func() {
			defer wg.Done()
			utils.MakeDir(utils.MakePath(flags.RootPath, flags.DirName), url.Host)
			// first
			os.Chdir(utils.MakePath(flags.RootPath, flags.DirName, url.Host))
			utils.DownloadPage(utils.MakePath(flags.RootPath, flags.DirName, url.Host), &url, client, flags.Recursive, flags.Level)
		}()
	}

	wg.Wait()
	fmt.Println("Successfully parsed")
}
