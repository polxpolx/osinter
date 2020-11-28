package secutrails

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func clientST(URL string, HttpMethod string, ApiKey string, json []byte) []byte {
	client := &http.Client{}

	req, err := http.NewRequest(HttpMethod, URL, bytes.NewBuffer(json))

	Checkerr(err)
	log.Printf("%s", json)
	req.Header.Add("APIKEY", ApiKey)

	resp, err := client.Do(req)
	Checkerr(err)
	defer resp.Body.Close()

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
