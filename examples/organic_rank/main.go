package main

import (
	"fmt"
	"log"
	"net/url"
	"netsurfer"
)

func main() {
	u, _ := url.Parse("https://github.com/ryonakao")
	rank, err := netsurfer.GetRank(u, "ryonakao", 2)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Rank is ", rank)
}
