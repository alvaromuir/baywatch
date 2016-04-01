package fetch

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/url"
	"os"
	"strings"
)

// Add your API Keys below
var bkuid = os.Getenv("BK_CMB_KEY")
var bksecretkey = os.Getenv("BK_CMB_SECRET")

const baseURL = "http://services.bluekai.com/Services/WS/"

// SignBKRequest returns a HMAC-SHA256 string for BlueKai JSON request
func SignBKRequest(method string, URL string, data string) string {
	URL = baseURL + URL
	stringToSign := method

	parsedURL, err := url.Parse(URL)
	if err != nil {
		panic(err)
	}

	stringToSign += parsedURL.Path

	qP := strings.Split(parsedURL.RawQuery, "&")
	if len(qP) > 0 {
		for qS := 0; qS < len(qP); qS++ {
			qP2 := strings.Split(qP[qS], "=")
			if len(qP2) > 1 {
				stringToSign += qP2[1]
			}
		}
	}

	if data != "" {
		stringToSign += data
	}

	h := hmac.New(sha256.New, []byte(bksecretkey))
	h.Write([]byte(stringToSign))
	s := base64.StdEncoding.EncodeToString(h.Sum(nil))
	u := url.QueryEscape(s)

	newURL := URL
	if strings.Contains(URL, "?") {
		newURL += "&"
	} else {
		newURL += "?"
	}

	newURL += "bkuid=" + bkuid + "&bksig=" + u
	print(newURL)
	return newURL
}
