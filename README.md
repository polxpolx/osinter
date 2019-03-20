# Osinter- A simple osint tool
Simple go program to do osint using different information passed as argument.

## Getting Started
`git clone https://github.com/reg0l/osinter/`

`go build osinter.go`


## Usage
```
Usage of ./osinter:

  -domain string
         the domain you want to look at
  -lookup string
        loolup against the targeted domain
  -reverse string
        reverse lookup against the domain
  -whois
        whois against the domain pass in argument
  -email
        email pass in argument for the pwned function
  -pwned
        check email passed in argument to Troy Hunt website https://haveibeenpwned.com/
 ```
 
 ## Example
 
Check if the email has been powned using famous webiste from Troy Hunt https://haveibeenpwned.com/
``./osinter -email theemailiwantocheck@emailprovider.com -pwned ``