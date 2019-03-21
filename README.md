# Osinter- A simple osint tool
Simple go program to do osint using different information passed as argument.

## Getting Started
`git clone https://github.com/reg0l/osinter/`

`go build osinter.go`


## Usage
```
Usage of ./osinter:
  -censys string
        call to censys.io api. Support the method account, data and search
  -censys-flatten
        Format the censys's search api results. Default is true (default true)
  -censys-index string
        Censys.io index for the search method.
  -censys-page int
        Censys.io pages number when search api invocked (default 1)
  -censys-query string
        Censys.io query for the search method.
  -domain string
         the domain you run the action against
  -email string
        email to check if this email has been powned
  -ip string
        ip use in different module
  -ipinfo
        check ip to ipinfo.io and return json
  -lookup
        domain lookup
  -powned
         check email passed in argument to Troy Hunt website https://haveibeenpwned.com/ and return json
  -reverse
        reverse lookup
  -reverse-ip
        reverse ip
  -whois
        domain whois
 ```
 
 ## Example
 
 ### HaveBeenPwned
Check if the email has been powned using famous webiste from Troy Hunt https://haveibeenpwned.com/

``./osinter -pwned -email myemailiwantocheck@emailprovider.com  ``

 ### Censys
 
 #### Censys Setup
 You first need to register to the website to get your api id and api secret : [signup](https://censys.io/register)
 
 Your Api credentials can be found in your account page than section: [API](https://censys.io/account/api)
 
 You need to setup local var environment variable to use censys.
 
 `export CENSYS_API_ID=REPLACE_WITH_YOUR_CENSYS_API_ID`
 
 `export CENSYS_API_SECRET=REPLACE_WITH_YOUR_CENSYS_API_SECRET`
 
  #### Censys Search Method
  
  
You will need to invocke th paramater `-censys search` .

Paramaters allowed 

* Required parameters - `-censys-index` can only be: `websites` or `certificates` or `ipv4`.
* Required parameters - `-censys-query`  the query to be executed. e.g `80.http.get.headers.server: nginx`.
* Optional paramters - `-censys-page` the page of the result set to be returned. 
* Optional parameters - `censys-fields` the fields you would like returned in the result set in "dot notation" e.g `location.country_code`
 
 Query example
 
 ` ./osinter -censys search -censys-query 'location.city: montpellier' -censys-index ipv4 -censys-page 2`

## TODO

[ ] Censys 

[ ] Shodan
