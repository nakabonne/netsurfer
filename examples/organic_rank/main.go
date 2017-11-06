package main

import (
	"fmt"
	"log"
	"net/url"
	"netsurfer"
)

func main() {
	u, _ := url.Parse("https://www.ruby-lang.org/ja/")
	rank, err := netsurfer.GetRank(u, "ruby", 1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Rank is ", rank)
}
