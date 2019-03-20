package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func prettifyJson(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err

}

func getResponse(srcURL, httpMethod string, header_useragent string, header_accept string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest(httpMethod, srcURL, nil)
	checkerr(err)
	req.Header.Add("User-Agent", header_useragent)
	req.Header.Add("Accept", header_accept)
	req.Header.Add("Accept-Language", "en-US,en;q=0.8")
	resp, err := client.Do(req)
	checkerr(err)
	defer resp.Body.Close()

	//check if the status code returned is wihtin the range 200-299. Fatal if not.
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fmt.Println("Something went wrong:", resp.StatusCode, "\n")
		bodybyte, err := ioutil.ReadAll(resp.Body)
		checkerr(err)
		fmt.Printf("%s", bodybyte)
		log.Fatal("Fatal error !")
	}
	bodybyte, err := ioutil.ReadAll(resp.Body)
	return bodybyte
}

func checkerr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
