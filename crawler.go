package netsurfer

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// SerpsURL returns some URLs displayed on the first page when you google search
func SerpsURL(word string) (urls []string, err error) {
	word = strings.Replace(word, " ", "+", -1)
	requestURL := "https://www.google.co.jp/search?rlz=1C5CHFA_enJP693JP693&q=" + string(word)
	baseURL, err := url.Parse(requestURL)
	if err != nil {
		return
	}
	doc, err := getDoc(baseURL)
	if err != nil {
		return
	}
	doc.Find(".r").Each(func(_ int, srg *goquery.Selection) {
		srg.Find("a").Each(func(_ int, s *goquery.Selection) {
			href, exists := s.Attr("href")
			if exists {
				reqURL, err := baseURL.Parse(href)
				if err == nil {
					urls = append(urls, reqURL.String())
				}
			}
		})
	})
	return
}

// GetHTML returns the response HTML when requesting for the given URL
func GetHTML(url string) (html string, err error) {
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(body)
	html = buf.String()
	return
}

// GetTitle returns the title tag of the HTML file indicated by the given URL
func GetTitle(baseURL string) (title string, err error) {
	var parsedURL *url.URL
	parsedURL, err = url.Parse(baseURL)
	if err != nil {
		return
	}
	doc, err := getDoc(parsedURL)
	if err != nil {
		return
	}
	doc.Find("title").Each(func(_ int, srg *goquery.Selection) {
		title = srg.Text()
	})
	return
}

func getDoc(u *url.URL) (doc *goquery.Document, err error) {
	res, err := http.Get(u.String())
	if err != nil {
		return
	}
	defer res.Body.Close()
	doc, err = goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}
	return
}
