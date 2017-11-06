package main

import (
	"fmt"
	"log"
	"netsurfer"
)

func main() {
	urls, err := netsurfer.OrganicSearch("ruby", 3)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Success!")
		for _, v := range urls {
			title, err := netsurfer.GetTitle(v.String())
			if err != nil {
				log.Panicln(err)
			}
			fmt.Println(title)
		}
	}
}
