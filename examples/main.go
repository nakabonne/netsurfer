package main

import (
	"fmt"
	"log"
	"netsurfer"
)

func main() {
	urls, err := netsurfer.SerpsURL("ruby")
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Success!")
		for _, v := range urls {
			title, err := netsurfer.GetTitle(v)
			if err != nil {
				log.Panicln(err)
			}
			fmt.Println(title)
		}
	}
}
