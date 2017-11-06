package main

import (
	"fmt"
	"log"

	"github.com/ryonakao/netsurfer"
)

func main() {
	html, err := netsurfer.GetHTML("https://qiita.com/ryonakao")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(html)
}
