package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/alvaromuir/baywatch/fetch"
	"github.com/alvaromuir/baywatch/serve"
)

func main() {

	flag.Usage = func() {
		fmt.Println("------------------------------------------------------------------------")
		fmt.Printf("Usage: $ %v [ server ] | [ callType, method, endPoint ] <opt: data > < opt: format >\n", strings.Split(os.Args[0], "/")[len(strings.Split(os.Args[0], "/"))-1])
		fmt.Println("------------------------------------------------------------------------")
		flag.PrintDefaults()
	}

	serverFlag := flag.Bool("server", false, "Run as server. Default: 'false'")
	callType := flag.String("callType", "empty", "API root. 'services' or 'taxonomy'")
	method := flag.String("method", "empty", "API verb, e.g. 'GET', 'POST', etc...")
	endPoint := flag.String("endPoint", "empty", "API endpoint resources'")
	data := flag.String("data", "", "Data to fetch or send, in JSON")
	format := flag.String("format", "table", "Determine on-screen display. 'table' (default), 'pretty' or 'raw'")

	flag.Parse()
	if *serverFlag == true {
		// run as webserver
		port := "8080"
		http.HandleFunc("/", serve.RootHandler)
		fmt.Println("server listening on http://localhost:" + port)
		fmt.Println("Ctrl-C to stop")
		log.Fatal(http.ListenAndServe("localhost:"+port, nil))
	} else {
		// run as utility

		if *callType == "empty" {
			flag.Usage()
			os.Exit(1)
		}
		if *method == "empty" {
			flag.Usage()
			os.Exit(1)
		}
		if *endPoint == "empty" {
			flag.Usage()
			os.Exit(1)
		}

		switch *endPoint {
		case "Ping":
			// returns status code from ping request
			resp, _, _ := fetch.PingRequest(
				fetch.SignRequest(*callType, *method, *endPoint, *data))
			pingResponse := fetch.PingResponse{Status: resp.StatusCode}
			rslt, err := json.Marshal(pingResponse)
			if err != nil {
				log.Fatal(err)
				fmt.Printf("error:%v\n", resp.StatusCode)
			}
			fmt.Println(string(rslt))
		case "categories":
			resp, rslt, err := fetch.TaxonomyBuyerRequest(
				*method, fetch.SignRequest(*callType, *method, *endPoint, *data),
				*data)

			if err != nil {
				log.Fatal(err)
			}

			if *format == "table" {
				fmt.Println("------------------")
				fmt.Fprintf(os.Stdout, "status: %d\n", resp.StatusCode)
				fmt.Fprintf(os.Stdout, "total count: %d\n", rslt.TotalResults)
				fmt.Println("------------------")
				if *data != "" {
					fmt.Printf("\n%s \t%13s \t%12s\n",
						"BKID", "REACH", "NAME")
					for _, item := range rslt.Items {
						fmt.Printf("#%d \t%-10d \t%s\n",
							item.ID, item.Stats.Reach, item.Name)
					}
				} else {
					fmt.Printf("\n%s \t%16s \t%s\n",
						"BKID", "PARENTID", "NAME")
					for _, item := range rslt.Items {
						fmt.Printf("#%d \t#%-10d \t%s\n",
							item.ID, item.ParentCategory.ID, item.Name)
					}
				}
			} else {
				if *format == "raw" {
					jsonData, err := json.Marshal(rslt)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println(string(jsonData))
				} else {
					jsonData, err := json.MarshalIndent(rslt, "", "    ")
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println(string(jsonData))
				}
			}
		case "sites":
			resp, rslt, err := fetch.SiteRequest(
				*method, fetch.SignRequest(*callType, *method, *endPoint, *data),
				*data)
			if err != nil {
				log.Fatal(err)
			}
			if *format == "table" {
				fmt.Println("------------------")
				fmt.Fprintf(os.Stdout, "status: %d\n", resp.StatusCode)
				fmt.Fprintf(os.Stdout, "total count: %d\n", rslt.TotalCount)
				fmt.Println("------------------")
				fmt.Printf("\n%s \t%s \t%20s\n",
					"BKID", "Last Updated", "NAME")
				for _, item := range rslt.Sites {
					fmt.Printf("#%d \t%s \t%s\n",
						item.ID, item.UpdatedAt, item.Name)
				}
			} else {
				if *format == "raw" {
					jsonData, err := json.Marshal(rslt)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println(string(jsonData))
				} else {
					jsonData, err := json.MarshalIndent(rslt, "", "    ")
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println(string(jsonData))
				}
			}
		}
	}
}
