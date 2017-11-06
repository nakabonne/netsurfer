package netsurfer

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func OrganicSearch(keyword string, depth int) (urls []*url.URL, err error) {
	var tmpURLs []*url.URL
	resultURL, err := serpsURL(keyword)
	if err != nil {
		return
	}
	for i := 0; i < depth; i++ {
		tmpURLs, err = organicURLs(resultURL)
		if err != nil {
			urls = nil
			return
		}
		urls = append(urls, tmpURLs...)
		if last(i, depth) {
			return
		}
		resultURL, err = nextPage(resultURL)
		if err != nil {
			urls = nil
			return
		}
	}
	return
}

func last(local int, last int) bool {
	return local == last-1
}

func GetRank(targetURL *url.URL, keyword string, depth int) (rank int, err error) {
	var urls []*url.URL
	resultURL, err := serpsURL(keyword)
	if err != nil {
		return
	}
	for i := 0; i < depth; i++ {
		urls, err = organicURLs(resultURL)
		for _, u := range urls {
			rank++
			if u == targetURL {
				return
			}
		}
		resultURL, err = nextPage(resultURL)
		if err != nil {
			rank = 0
			return
		}
	}
	return 0, errors.New("That page is out of rank")
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

// serpsURL returns some URLs displayed on the first page when you google search
func serpsURL(word string) (serpsURL *url.URL, err error) {
	word = strings.Replace(word, " ", "+", -1)
	requestURL := "https://www.google.co.jp/search?rlz=1C5CHFA_enJP693JP693&q=" + string(word)
	serpsURL, err = url.Parse(requestURL)
	return
}

func organicURLs(reqURL *url.URL) (urls []*url.URL, err error) {
	doc, err := getDoc(reqURL)
	if err != nil {
		return
	}
	doc.Find(".r").Each(func(_ int, srg *goquery.Selection) {
		srg.Find("a").Each(func(_ int, s *goquery.Selection) {
			href, exists := s.Attr("href")
			if exists {
				reqURL, err := reqURL.Parse(href)
				if err != nil {
					urls = nil
					return
				}
				urls = append(urls, reqURL)
			}
		})
	})
	return
}

func nextPage(baseURL *url.URL) (nextURL *url.URL, err error) {
	doc, err := getDoc(baseURL)
	doc.Find("#nav").Each(func(_ int, table *goquery.Selection) {
		table.Find("tbody").Each(func(_ int, trs *goquery.Selection) {
			trs.Find("tr").Each(func(_ int, tds *goquery.Selection) {
				tds.Find("td").Each(func(_ int, srg *goquery.Selection) {
					srg.Find("a").Each(func(_ int, s *goquery.Selection) {
						href, exists := s.Attr("href")
						if exists {
							nextURL, err = baseURL.Parse(href)
						}
					})
				})
			})
		})
	})
	fmt.Println("次のURLは", nextURL)
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
