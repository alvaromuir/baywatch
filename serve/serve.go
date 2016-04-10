package serve

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/alvaromuir/baywatch/fetch"
)

func param(s []string, i int) string {
	defer func() {
		if err := recover(); err != nil {
			// fmt.Print(err)
			return
		}
	}()
	return s[i]
}

func setHeader(w http.ResponseWriter, errNumber int) {
	w.WriteHeader(errNumber)
	w.Header().Set("Server", "vz media science: baywatch")
	w.Header().Set("Content-Type", "application/json")
}

// RootHandler returns some stuff
func RootHandler(w http.ResponseWriter, r *http.Request) {
	urlArgs := strings.Split(r.URL.Path, "/")
	callType := param(urlArgs, 1)
	method := param(urlArgs, 2)
	endPoint := param(urlArgs, 3)
	data := r.URL.RawQuery

	switch endPoint {
	case "Ping":
		resp, _, _ := fetch.PingRequest(
			fetch.SignRequest(callType, method, endPoint, data))
		setHeader(w, resp.StatusCode)
		// manual, because Ping call returns no body ...
		pingResponse := fetch.PingResponse{Status: resp.StatusCode}
		rslt, err := json.Marshal(pingResponse)
		if err != nil {
			log.Fatal(err)
			fmt.Fprintf(w, "error: %v", resp.StatusCode)
		}
		w.Write(rslt)

	case "categories":
		resp, rslt, err := fetch.TaxonomyBuyerRequest(
			method, fetch.SignRequest(callType, method, endPoint, data),
			data)
		if err != nil {
			log.Fatal(err)
			setHeader(w, resp.StatusCode)
			fmt.Fprintf(w, "{status: %d\n}", resp.StatusCode)
		}
		setHeader(w, resp.StatusCode)
		jsonData, err := json.MarshalIndent(rslt, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, string(jsonData))
	case "sites":
		resp, rslt, err := fetch.SiteRequest(
			method, fetch.SignRequest(callType, method, endPoint, data),
			data)
		if err != nil {
			log.Fatal(err)
			setHeader(w, resp.StatusCode)
			fmt.Fprintf(w, "{status: %d\n}", resp.StatusCode)
		}
		setHeader(w, resp.StatusCode)
		jsonData, err := json.MarshalIndent(rslt, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, string(jsonData))
	}
}
