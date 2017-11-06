package main

import (
	"fmt"
	"log"
	"netsurfer"
)

func main() {
	// Obtain the URL of the organic page
	urls, err := netsurfer.OrganicSearch("ruby", 3)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Success!")
	for _, v := range urls {
		// Retrieve the title
		title, err := netsurfer.GetTitle(v.String())
		if err != nil {
			log.Panicln(err)
		}
		fmt.Println(title)
	}

}