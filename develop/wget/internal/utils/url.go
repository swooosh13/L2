package utils

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/url"
	"strings"
)

// Strings To url format

func StringToUrl(s string) (*url.URL, error) {
	url, err := url.Parse(s)
	if err != nil {
		return nil, err
	}

	if url.Scheme == "" {
		url.Scheme = DEFAULT_URL_SCHEME
	}

	_, ok := UrlSchemes[url.Scheme]
	if !ok {
		return nil, fmt.Errorf("url scheme not supported: %s", url.Scheme)
	}

	return url, nil
}

func StringsToUrls(us []string) ([]url.URL, error) {
	urls := make([]url.URL, 0)
	for _, u := range us {
		url, err := StringToUrl(u)
		if err != nil {
			return nil, err
		}
		urls = append(urls, *url)
	}
	return urls, nil
}

func PageUrls(r io.Reader, url url.URL) ([]url.URL, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	var urls []string
	doc.Find("a").Each(func(i int, selection *goquery.Selection) {
		linkTag := selection
		link, _ := linkTag.Attr("href")

		if strings.HasPrefix(link, "/") && link != url.Path {
			urls = append(urls, url.Scheme+"://"+url.Host+link)
		} else if strings.HasPrefix(link, url.Scheme+"://"+url.Host) && link != url.String() {
			urls = append(urls, link)
		}
	})

	return StringsToUrls(urls)
}
