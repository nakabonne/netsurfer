package main

import (
	"fmt"
	"netsurfer"
)

func main() {
	urls, err := netsurfer.SerpsURL("ruby")
	if err != nil {
		fmt.Println("erorr!", err)
	} else {
		fmt.Println("Success!")
		for _, v := range urls {
			fmt.Println(v)
		}
	}
}
