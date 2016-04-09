package fetch

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Request generates the method request
func Request(method string, URL string, data string) string {
	// cj, _ := cookiejar.New(nil)
	client := &http.Client{}
	req, _ := http.NewRequest(method, URL, strings.NewReader(data))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, os.Args[0]+": %v", err)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, os.Args[0]+": %v", err)
		os.Exit(1)
	}

	if len(body) < 1 {
		return "Status: " + strconv.Itoa(res.StatusCode)
	}
	return bytes.NewBuffer(body).String()
}
