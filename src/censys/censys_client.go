package censys

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func ClientCensys(ApiId string, ApiSecret string, url string, timeout int, HttpMethod string, HeaderUserAgent string, HeaderAccept string, json []byte) []byte {
	client := &http.Client{}

	req, err := http.NewRequest(HttpMethod, url, bytes.NewBuffer(json))

	Checkerr(err)
	log.Printf("%s", json)
	//setup cleaned http header
	req.Header.Add("User-Agent", HeaderUserAgent)
	req.Header.Add("Accept", HeaderAccept)
	req.Header.Add("Accept-Language", "en-US,en;q=0.8")

	// setup http auth
	req.SetBasicAuth(ApiId, ApiSecret)

	// run the client request
	resp, err := client.Do(req)
	Checkerr(err)
	defer resp.Body.Close()

	// check query could not be parsed
	if resp.StatusCode == 400 {
		log.Fatalln("Query could not be parsed: ", resp.StatusCode)
	}

	// page not found
	if resp.StatusCode == 404 {
		log.Fatal("Page not found: ", resp.StatusCode)
	}

	// check authentication issues
	if resp.StatusCode == 403 {
		log.Fatalln("Unauthorized. You must authenticate with a valid API ID and secret: ", resp.StatusCode)
	}

	// check if rate limit has been exceeded
	if resp.StatusCode == 429 {
		log.Fatalln("Rate limit exceeded: ", resp.StatusCode)
	}

	// check if unknown error happened
	if resp.StatusCode == 500 {
		log.Fatalln("Unknown error occurred: ", resp.StatusCode)
	}

	//check if the status code returned is different of 200. Fatal if not.
	if resp.StatusCode != 200 {
		fmt.Println("Something went wrong: ", resp.StatusCode)
		bodybyte, err := ioutil.ReadAll(resp.Body)
		Checkerr(err)
		log.Printf("%s", bodybyte)
		log.Fatal("Fatal error !")
	}

	bodybyte, err := ioutil.ReadAll(resp.Body)
	return bodybyte
}

func Checkerr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
