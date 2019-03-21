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
const HavebeenpwnedHeaderAccept string = "application/vnd.haveibeenpwned.v2+json"
const HavebeenpwnedHeaderUserAgent string = "OSINTER-Osint-Tools-For-Defense-Team-OSINTER"

//generic variables
const generic_header_user_agent string = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36"
const generic_header_accept string = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"

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
	censysin := flag.String("censys", "", "call to censys.io api. Support the method account, data and search")
	censysquery := flag.String("censys-query", "", "Censys.io query for the search method.")
	censysindex := flag.String("censys-index", "", "Censys.io index for the search method.")
	censyspage := flag.Int("censys-page", 1, "Censys.io pages number when search api invocked")
	censysflatten := flag.Bool("censys-flatten", true, "Format the censys's search api results. Default is true")
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
	if *censysin != "" {

		// function to pick env variables to be consumed by Censys client
		CensysApiId := os.Getenv("CENSYS_API_ID")
		CensysApiSecret := os.Getenv("CENSYS_API_SECRET")

		// we verify if the env variable has been set and are not empty.
		if CensysApiId != "" && CensysApiSecret != "" {
			switch *censysin {
			case "account":
				// Call Censys account method
				body := censys.ClientCensys(CensysApiId, CensysApiSecret, censys.CensysUrlAccount, 30, "GET", censys.CensysHeaderUserAgent, censys.CensysHeaderAccept, nil)

				// Prettify print of the json
				utils.PrettifyPrint(body)

			case "data":
				// Call Censys data method
				body := censys.ClientCensys(CensysApiId, CensysApiSecret, censys.CensysUrlData, 30, "GET", censys.CensysHeaderUserAgent, censys.CensysHeaderAccept, nil)

				// Prettify print of the json
				utils.PrettifyPrint(body)

			case "search":
				if *censysquery != "" && *censysindex != "" {
					// init CensysJson Struct
					search := censys.CensysJson{}

					// fil CensysJson struct. See censys_model.go for more info
					search.Query = *censysquery
					search.Pages = *censyspage
					search.Flatten = *censysflatten

					// Marshalized search element to be passed to CensysClient
					jsonMarshaled := utils.Marshallizer(search)

					// TODO manage exception using more elegant manner
					switch *censysindex {
					case "ipv4":
						// prepare url string concat CensysSearchURL and CensysIndex & check if censys-index argument is valid. If not FATAL
						urlConcat := censys.CensysUrlSearch + "/" + censys.CensysIndexIPV4
						log.Println("Prepare call to: ", urlConcat)
						// Call Censys search method
						body := censys.ClientCensys(CensysApiId, CensysApiSecret, urlConcat, 30, "POST", censys.CensysHeaderUserAgent, censys.CensysHeaderAccept, jsonMarshaled)

						// Prettify print of the json
						utils.PrettifyPrint(body)

					case "certificates":
						// prepare url string concat CensysSearchURL and CensysIndex & check if censys-index argument is valid. If not FATAL
						urlConcat := censys.CensysUrlSearch + "/" + censys.CensysIndexCertificate

						log.Println("Prepare call to: ", urlConcat)

						// Call Censys search method
						body := censys.ClientCensys(CensysApiId, CensysApiSecret, urlConcat, 30, "POST", censys.CensysHeaderUserAgent, censys.CensysHeaderAccept, jsonMarshaled)

						// Prettify print of the json
						utils.PrettifyPrint(body)

					case "websites":
						// prepare url string concat CensysSearchURL and CensysIndex & check if censys-index argument is valid. If not FATAL
						urlConcat := censys.CensysUrlSearch + "/" + censys.CensysIndexWebsites

						log.Println("Prepare call to: ", urlConcat)
						// Call Censys search method
						body := censys.ClientCensys(CensysApiId, CensysApiSecret, urlConcat, 30, "POST", censys.CensysHeaderUserAgent, censys.CensysHeaderAccept, jsonMarshaled)

						// Prettify print of the json
						utils.PrettifyPrint(body)

					default:
						log.Fatalln(" A valid censys-index argument need to be pass. We're accepting: websites, ipv4 or certificates.")
					}

				} else {
					log.Fatalln("The argument can not be empty with the censys's search method. Please provide valid argument e.g: 80.http.get.headers.server: nginx")
				}

			default:
				log.Fatalln("Please provide a valid method for censys. We're supporting: data and account.")
			}

		} else {
			log.Fatalln("No Cencys api id or key provided in environment variable.")
		}
	}
}

func getWhois(domain string) {
	srcUrl := fmt.Sprintf("https://api.hackertarget.com/whois/?q=%s", domain)
	body := utils.GetResponse(srcUrl, "GET", generic_header_user_agent, generic_header_accept)
	fmt.Println(string(body))
}

func getReverseDNS(domain string) {
	srcUrl := fmt.Sprintf("https://api.hackertarget.com/reversedns/?q=%s", domain)
	body := utils.GetResponse(srcUrl, "GET", generic_header_user_agent, generic_header_accept)
	fmt.Println(string(body))
}

func getDNSLookup(domain string) {
	srcUrl := fmt.Sprintf("https://api.hackertarget.com/dnslookup/?q=%s", domain)
	body := utils.GetResponse(srcUrl, "GET", generic_header_user_agent, generic_header_accept)
	fmt.Println(string(body))
}

func getReverseIp(domain string) {
	srcUrl := fmt.Sprintf("https://api.hackertarget.com/reverseiplookup/?q=%s", domain)
	body := utils.GetResponse(srcUrl, "GET", generic_header_user_agent, generic_header_accept)
	fmt.Println(string(body))
}

func getHaveBeenPowned(email string) {
	srcUrl := fmt.Sprintf("https://haveibeenpwned.com/api/v2/breachedaccount/%s", email)
	body := utils.GetResponse(srcUrl, "GET", HavebeenpwnedHeaderUserAgent, HavebeenpwnedHeaderAccept)
	prettify_body, _ := utils.PrettifyJson(body)
	fmt.Printf("%s", prettify_body)
}

func getIpInfo(ipcheck string) {
	srcUrl := fmt.Sprintf("https://ipinfo.io/%s/json", ipcheck)
	body := utils.GetResponse(srcUrl, "GET", generic_header_user_agent, "application/json")
	// Prettify print of the json
	utils.PrettifyPrint(body)
}
