package main

import (
	"os"

	"github.com/alvaromuir/baywatch/fetch"
)

func main() {
	verb := os.Args[1]
	endPoint := os.Args[2]
	var data string

	if len(os.Args) > 3 {
		if len(os.Args[3]) > 1 {
			data = os.Args[3]
		}
	}

	newURL := fetch.SignBKRequest(verb, endPoint, data)
	fetch.Request(verb, newURL, data)
}
