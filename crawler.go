package netsurfer

import (
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func SerpsURL(word string) (urls []string, err error) {

	log.Println("検索ワード：", word)
	word = strings.Replace(word, " ", "+", -1)
	requestURL := "https://www.google.co.jp/search?rlz=1C5CHFA_enJP693JP693&q=" + string(word)
	log.Println("検索URL：", requestURL)

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
