# netsurfer

netsurfer is a very lightweight scraping framework

# What can be done

- You can know the ranking of the web site.
- You can know the URL displayed when searching with the google engine.
- You can get HTML.

# Installation

```
$ go get -u github.com/ryonakao/netsurfer
```

# Usage

If you want to know more, please read [examples](https://github.com/ryonakao/netsurfer/tree/master/examples).

```go
import (
	"fmt"
	"net/url"
	"github.com/ryonakao/netsurfer"
)
// If you want to know the rank of a page, please write the following code.
u, _ := url.Parse("https://qiita.com/ryonakao")
rank, _ := netsurfer.GetRank(u, "ryonakao", 2)
fmt.Println("Rank is ", rank)

// If you want to know the search result URL, please write following
urls, _ := netsurfer.OrganicSearch("ruby", 3)

for _, url := range urls {
	// If you want to know the title, please write the following
	title, _ := netsurfer.GetTitle(url.String())
	fmt.Println("Title is ", title)
}
```

# dependencies

- [goquery](https://github.com/PuerkitoBio/goquery)

# License

`netsurfer` source code is available under the MIT [License](https://github.com/ryonakao/netsurfer/blob/master/LICENSE).