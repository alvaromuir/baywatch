package fetch

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Request generates the method request
func Request(method string, URL string, data string) {
	// cj, _ := cookiejar.New(nil)
	client := &http.Client{}
	req, _ := http.NewRequest(method, URL, strings.NewReader(data))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/json")
	res, _ := client.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(bytes.NewBuffer(body).String())
}
