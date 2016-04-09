package fetch

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var env = godotenv.Load(".env")

// Add your API Keys below
var bkuid = os.Getenv("BK_KEY")              // your BlueKai API KEY
var bksecretkey = os.Getenv("BK_SECRET")     // your BlueKai API Secret
var bkpartnerid = os.Getenv("BK_PARTNER_ID") // your BlueKai partner ID

var resources = map[string]string{
	"services": "http://services.bluekai.com/Services/WS/",
	"taxonomy": "https://taxonomy.bluekai.com/taxonomy/",
}

// SignRequest returns a HMAC-SHA256 string for BlueKai JSON request
func SignRequest(baseURL string, method string, endPoint string, data string) string {
	var URL string
	if baseURL == "services" {
		URL = resources["services"] + endPoint
	}

	if baseURL == "taxonomy" {
		URL = resources["taxonomy"] + endPoint + "?" + data
		if strings.HasPrefix(data, "parentCategory") {
			URL += "&view=BUYER"
		} else {
			URL += "&view=OWNER"
		}
		URL += "&showReach=true&countryCode=US&partner.id=" + bkpartnerid
	}

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
	// fmt.Println(newURL)
	return newURL
}
