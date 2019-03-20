package main

import (
	"flag"
	"fmt"
	"github.com/reg0l/osinter/src/censys"
	"github.com/reg0l/osinter/src/utils"
	"log"
	"os"
)

//havebeenpwned's constant variables
const havebeenpowned_header_accept string = "application/vnd.haveibeenpwned.v2+json"
const havebeenpowned_header_user_agent string = "Osint-Tools-For-Defense-Team-OSINTER"

//generic variables
const generic_header_user_agent string = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36"
const generic_header_accept string = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"

//censys' constant variables
const censysaccounturl string = "https://www.censys.io/api/v1/account"
const censysacceptheader string = "application/json, */8"
const censysheaderuseragent string = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36"

func main() {

	domain := flag.String("domain", "", " the domain you run the action against")
	ip := flag.String("ip", "", "ip use in different module")
	whois := flag.Bool("whois", false, "domain whois")
	reverse := flag.Bool("reverse", false, "reverse lookup")
	reverseIp := flag.Bool("reverse-ip", false, "reverse ip")
	lookup := flag.Bool("lookup", false, "domain lookup")
	powned := flag.Bool("powned", false, " check email passed in argument to Troy Hunt website https://haveibeenpwned.com/ and return json")
	email := flag.String("email", "", "email to check if this email has been powned")
	ipinfo := flag.Bool("ipinfo", false, "check ip to ipinfo.io and return json")
	censysaccount := flag.Bool("censys-account", false, "call to censys.io api to get your account information")
	flag.Parse()

	if *whois && *domain != "" {
		getWhois(*domain)
	}

	if *lookup && *domain != "" {
		getDNSLookup(*domain)
	}

	if *reverse && *domain != "" {
		getReverseDNS(*domain)
	}

	if *reverseIp && *ip != "" {
		getReverseIp(*ip)
	}
	if *ipinfo && *ip != "" {
		getIpInfo(*ip)
	}
	if *powned && *email != "" {
		getHaveBeenPowned(*email)
	}
	if *censysaccount {

		// function to pick env variables to be consumed by Censys client
		apiidenv := os.Getenv("CENSYS_API_ID")
		apisecretenv := os.Getenv("CENSYS_API_SECRET")

		// we verify if the env variable has been set and are not empty.
		if apiidenv != "" && apisecretenv != "" {
			body := censys.ClientCensys(apiidenv, apisecretenv, censysaccounturl, 30, "GET", censysheaderuseragent, censysacceptheader)
			log.Printf("%s", body)
		} else {
			log.Fatalln("No cencys api id or key provided in environment variable.")
		}
	}
}

func getWhois(domain string) {
	srcUrl := fmt.Sprintf("https://api.hackertarget.com/whois/?q=%s", domain)
	body := utils.GetResponse(srcUrl, "GET", generic_header_user_agent, generic_header_accept)
	log.Println(string(body))
}

func getReverseDNS(domain string) {
	srcUrl := fmt.Sprintf("https://api.hackertarget.com/reversedns/?q=%s", domain)
	body := utils.GetResponse(srcUrl, "GET", generic_header_user_agent, generic_header_accept)
	log.Println(string(body))
}

func getDNSLookup(domain string) {
	srcUrl := fmt.Sprintf("https://api.hackertarget.com/dnslookup/?q=%s", domain)
	body := utils.GetResponse(srcUrl, "GET", generic_header_user_agent, generic_header_accept)
	log.Println(string(body))
}

func getReverseIp(domain string) {
	srcUrl := fmt.Sprintf("https://api.hackertarget.com/reverseiplookup/?q=%s", domain)
	body := utils.GetResponse(srcUrl, "GET", generic_header_user_agent, generic_header_accept)
	log.Println(string(body))
}

func getHaveBeenPowned(email string) {
	srcUrl := fmt.Sprintf("https://haveibeenpwned.com/api/v2/breachedaccount/%s", email)
	body := utils.GetResponse(srcUrl, "GET", havebeenpowned_header_user_agent, havebeenpowned_header_accept)
	prettify_body, _ := utils.PrettifyJson(body)
	log.Printf("%s", prettify_body)
}

func getIpInfo(ipcheck string) {
	srcUrl := fmt.Sprintf("https://ipinfo.io/%s/json", ipcheck)
	body := utils.GetResponse(srcUrl, "GET", generic_header_user_agent, "application/json")
	prettifyBody, _ := utils.PrettifyJson(body)
	log.Printf("%s", prettifyBody)
}
