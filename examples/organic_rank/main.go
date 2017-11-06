package main

import (
	"fmt"
	"net/url"

	"github.com/ryonakao/netsurfer"
)

func main() {
	u, _ := url.Parse("https://qiita.com/ryonakao")
	rank, _ := netsurfer.GetRank(u, "ryonakao", 2)
	fmt.Println("Rank is ", rank)
}
