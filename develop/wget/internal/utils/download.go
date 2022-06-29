package utils

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var seen = map[string]struct{}{}

// path - root site path

func DownloadPage(path string, url *url.URL, client *http.Client, recursive bool, level int) {
	var fileName string
	if url.Path == "/" {
		fileName = "index"
	} else {
		segments := strings.Split(url.Path, "/")
		fileName = segments[len(segments)-1]
	}

	file, err := os.Create(fileName + ".html")
	if err != nil {
		fmt.Println("error create file: ", err)
		os.Exit(1)
	}
	defer file.Close()

	response, err := client.Get(url.String())
	if err != nil {
		log.Fatalln("error request: ", err)
	}
	defer response.Body.Close()

	fmt.Printf("Request Status: %s\n\n", response.Status)

	var buf bytes.Buffer
	tee := io.TeeReader(response.Body, &buf)

	_, err = io.Copy(file, tee)
	if err != nil {
		log.Fatalln("error copy content: ", err)
	}

	if recursive && level > 0 {
		urls, err := PageUrls(&buf, *url)
		if err != nil {
			log.Fatalln(err)
		}

		for _, u := range urls {
			if _, ok := seen[u.String()]; ok {
				continue
			}
			seen[u.String()] = struct{}{}

			p := strings.TrimSuffix(url.Path, "/")
			segments := strings.Split(p, "/")
			segments = segments[:len(segments)-1]
			subPath := strings.Join(segments, "/")

			MakeDir(path, subPath)
			os.Chdir(path + subPath)

			DownloadPage(path, &u, client, true, level-1)
		}
	}
}
