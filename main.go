package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/alvaromuir/baywatch/fetch"
	"github.com/alvaromuir/baywatch/serve"
)

func main() {
	// fmt.Println(len(os.Args))
	if len(os.Args) > 2 {
		callType := os.Args[1]
		method := os.Args[2]
		endPoint := os.Args[3]

		var data string

		if len(os.Args) > 4 {
			if len(os.Args[4]) > 1 {
				data = os.Args[4]
			}
		}
		switch callType {
		case "services":
			// ex: .baywatch services GET sites
			rslt := fetch.Request(method, fetch.SignServicesRequest(method, endPoint, data), data)
			fmt.Println(rslt)
		case "taxonomy":
			// ex: .baywatch taxonomy GET categories
			// ex: .baywatch taxonomy GET categories "parentCategory.id=12345"
			rslt := fetch.Request(method, fetch.SignTaxonomyRequest(method, endPoint, data), data)
			fmt.Println(rslt)
		}
	} else {
		// ex: http://localhost:8080/taxonomy/GET/categories
		// ex: http://localhost:8080/taxonomy/GET/categories?parentCategory.id=399598
		if os.Args[1] == "server" {
			port := "8080"
			http.HandleFunc("/", serve.RootHandler)
			fmt.Println("server listening on http://localhost:" + port)
			fmt.Println("Ctrl-C to stop")
			log.Fatal(http.ListenAndServe("localhost:"+port, nil))
		} else {
			fmt.Println("error . . . ")
		}
	}
}
