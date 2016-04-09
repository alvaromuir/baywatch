package serve

import (
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

// RootHandler returns some stuff
func RootHandler(w http.ResponseWriter, r *http.Request) {
	urlArgs := strings.Split(r.URL.Path, "/")
	callType := param(urlArgs, 1)
	method := param(urlArgs, 2)
	endPoint := param(urlArgs, 3)
	data := r.URL.RawQuery

	var rslt string
	switch callType {
	case "services":
		// ex: .baywatch services GET sites
		rslt = fetch.Request(method, fetch.SignServicesRequest(method, endPoint, data), data)
	case "taxonomy":
		// ex: .baywatch taxonomy GET categories "parentCategory.id=12345"
		rslt = fetch.Request(method, fetch.SignTaxonomyRequest(method, endPoint, data), data)
	}
	w.Header().Set("Server", "vz media science: baywatch")
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(rslt))
}
