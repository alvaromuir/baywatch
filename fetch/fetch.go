package fetch

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// PingRequest generates the method request
func PingRequest(URL string) (*http.Response, *PingResponse, error) {
	method := "GET"
	data := ""
	resp, err := CoreRequest(method, URL, data)
	if err != nil {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, os.Args[0]+": %v", err)
		os.Exit(1)
	}
	var rslt PingResponse
	if err := json.NewDecoder(resp.Body).Decode(&rslt); err != nil {
		resp.Body.Close()
		return resp, nil, err
	}
	resp.Body.Close()
	return resp, &rslt, nil
}

// TaxonomyBuyerRequest returns DMP dites in JSON format
func TaxonomyBuyerRequest(method string, URL string, data string) (*http.Response, *BuyerViewCategoryResult, error) {
	resp, err := CoreRequest(method, URL, data)
	if err != nil {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, os.Args[0]+": %v", err)
		os.Exit(1)
	}
	var rslt BuyerViewCategoryResult
	if err := json.NewDecoder(resp.Body).Decode(&rslt); err != nil {
		resp.Body.Close()
		return resp, nil, err
	}
	return resp, &rslt, nil
}

// SiteRequest returns DMP dites in JSON format
func SiteRequest(method string, URL string, data string) (*http.Response, *SiteResult, error) {
	resp, err := CoreRequest(method, URL, data)
	if err != nil {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, os.Args[0]+": %v", err)
		os.Exit(1)
	}
	var rslt SiteResult
	if err := json.NewDecoder(resp.Body).Decode(&rslt); err != nil {
		resp.Body.Close()
		return resp, nil, err
	}
	return resp, &rslt, nil
}

// CoreRequest is a bare-bones JSON request, essential for all other API calls
func CoreRequest(method string, URL string, data string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, URL, strings.NewReader(data))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		resp.Body.Close()

		fmt.Fprintf(os.Stderr, os.Args[0]+": %v", err)
		os.Exit(1)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("API request failed: %s", resp.Status)
	}
	return resp, nil
}
