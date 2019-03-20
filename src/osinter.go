package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const havebeenpowned_header_user_agent string = "Osint-Tools-For-Defense-Team-OSINTER"
const generic_header_user_agent string = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36"
const generic_header_accept string = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"
const havebeenpowned_header_accept string = "application/vnd.haveibeenpwned.v2+json"

func main() {

	domainIn := flag.String("domain", "", " the domain you run the action against")
	ipIn := flag.String("ip", "", "ip for the reverse ip")
	whois := flag.Bool("whois", false, "domain whois")
	reverse := flag.Bool("reverse", false, "reverse lookup")
	reverseIp := flag.Bool("reverse-ip", false, "reverse ip")
	lookup := flag.Bool("lookup", false, "domain lookup")
	powned := flag.Bool("powned", false, "email been powned")
	email := flag.String("email", "", "email to check if this email has been powned")
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

	if *reverseIp {
		getReverseIp(*ipIn)
	}

	if *powned {
		getHaveBeenPowned(*email)
	} else {
		log.Fatalln("No valid parameter passed")
	}
}

func getWhois(domain string) {
	srcUrl := fmt.Sprintf("https://api.hackertarget.com/whois/?q=%s", domain)
	body := getResponse(srcUrl, "GET", generic_header_user_agent, generic_header_accept)
	fmt.Println(string(body))
}

func getReverseDNS(domain string) {
	srcUrl := fmt.Sprintf("https://api.hackertarget.com/reversedns/?q=%s", domain)
	body := getResponse(srcUrl, "GET", generic_header_user_agent, generic_header_accept)
	fmt.Println(string(body))
}

func getDNSLookup(domain string) {
	srcUrl := fmt.Sprintf("https://api.hackertarget.com/dnslookup/?q=%s", domain)
	body := getResponse(srcUrl, "GET", generic_header_user_agent, generic_header_accept)
	fmt.Println(string(body))
}

func getReverseIp(domain string) {
	srcUrl := fmt.Sprintf("https://api.hackertarget.com/reverseiplookup/?q=%s", domain)
	body := getResponse(srcUrl, "GET", generic_header_user_agent, generic_header_accept)
	fmt.Println(string(body))
}

func getHaveBeenPowned(email string) {
	srcUrl := fmt.Sprintf("https://haveibeenpwned.com/api/v2/breachedaccount/%s", email)
	body := getResponse(srcUrl, "GET", havebeenpowned_header_user_agent, havebeenpowned_header_accept)
	//fmt.Printf("%s", body)
	prettify_body, _ := prettifyJson(body)
	fmt.Printf("%s", prettify_body)
}

// TODO Google connector
//func getGooglePage(query string, page_no int){
//
//}

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

	//check if the status code returned is wihtin the range 200-299. Fatat if not.
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
