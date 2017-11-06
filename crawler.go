package netsurfer

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// OrganicSearch returns the page URL of organic search.
// If you want to search up to 3 pages, please give 3 to depth.
func OrganicSearch(keyword string, depth int) (urls []*url.URL, err error) {
	resultURL, err := serpsURL(keyword)
	if err != nil {
		return
	}
	serpsPages := []*url.URL{resultURL}
	tmpURLs := []*url.URL{}

	pages, err := getSERPsURLs(resultURL, depth)
	if err != nil {
		return
	}
	serpsPages = append(serpsPages, pages...)
	for _, page := range serpsPages {
		tmpURLs, err = organicURLs(page)
		if err != nil {
			return
		}
		urls = append(urls, tmpURLs...)
	}
	return
}

// GetRank returns the rank of the specified page when you search with the specified keyword.
// If you want to search up to 3 pages, please give 3 to depth.
func GetRank(targetURL *url.URL, keyword string, depth int) (rank int, err error) {
	resultURL, err := serpsURL(keyword)
	if err != nil {
		return
	}
	urls := []*url.URL{}
	serpsPages := []*url.URL{resultURL}
	pages, err := getSERPsURLs(resultURL, depth)
	if err != nil {
		return
	}
	serpsPages = append(serpsPages, pages...)
	for _, page := range serpsPages {
		urls, err = organicURLs(page)
		for _, u := range urls {
			query, _ := url.ParseQuery(u.RawQuery)
			rank++
			if sameURL(query["q"][0], targetURL.String()) {
				return
			}
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

func getSERPsURLs(baseURL *url.URL, depth int) (pages []*url.URL, err error) {
	doc, err := getDoc(baseURL)
	i := 0
	doc.Find("#nav").Each(func(_ int, table *goquery.Selection) {
		table.Find("tbody").Each(func(_ int, trs *goquery.Selection) {
			trs.Find("tr").Each(func(_ int, tds *goquery.Selection) {
				tds.Find("td").Each(func(_ int, srg *goquery.Selection) {
					srg.Find("a").Each(func(_ int, s *goquery.Selection) {
						if i >= depth-1 {
							return
						}
						href, exists := s.Attr("href")
						if exists {
							nextURL, _ := baseURL.Parse(href)
							pages = append(pages, nextURL)
							i++
						} else {
							err = errors.New("failed to retrieve the search result page")
							return
						}
					})
				})
			})
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

func last(local int, last int) bool {
	return local == last-1
}

func sameURL(urlA string, urlB string) bool {
	return urlA == urlB
}
