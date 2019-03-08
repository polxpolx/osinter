package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//getWhois("google.com")
	//getReverseDNS("google.com")
	//getDNSLookup("google.com")

	domainIn := flag.String("domain", "", " the domain you want to look at")
	whois := flag.Bool("whois", false, "whois against the domain pass in argument")
	reverse := flag.Bool("reverse", false, "reverse lookup against the domain")
	lookup := flag.Bool("lookup", false, "loolup against the targeted domain")
	flag.Parse()

	if *whois {
		getWhois(*domainIn)
	}
	if *lookup {
		getDNSLookup(*domainIn)
	}
	if *reverse {
		getReverseDNS(*domainIn)
	}
}

func getWhois(domain string) {
	srcUrl := fmt.Sprintf("https://api.hackertarget.com/whois/?q=%s", domain)
	body := getResponse(srcUrl, "GET")
	fmt.Println(string(body))
}

func getReverseDNS(domain string) {
	srcUrl := fmt.Sprintf("https://api.hackertarget.com/reversedns/?q=%s", domain)
	body := getResponse(srcUrl, "GET")
	fmt.Println(string(body))
}

func getDNSLookup(domain string) {
	srcUrl := fmt.Sprintf("https://api.hackertarget.com/dnslookup/?q=%s", domain)
	body := getResponse(srcUrl, "GET")
	fmt.Println(string(body))
}

func getResponse(srcURL, httpMethod string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest(httpMethod, srcURL, nil)
	checkerr(err)
	resp, err := client.Do(req)
	checkerr(err)
	defer resp.Body.Close()

	bodybyte, err := ioutil.ReadAll(resp.Body)
	checkerr(err)
	return bodybyte
}

func checkerr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
