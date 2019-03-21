package censys

const (
	// https://censys.io/api/v1/docs/account
	CensysUrlAccount string = "https://www.censys.io/api/v1/account"

	// https://censys.io/api/v1/docs/data
	CensysUrlData string = "https://www.censys.io/api/v1/data"

	// https://censys.io/api/v1/docs/search
	CensysUrlSearch string = "https://www.censys.io/api/v1/search"

	// https://censys.io/api/v1/docs/view
	CensysUrlView string = "https://www.censys.io/api/v1/view/"

	// https://censys.io/api/v1/docs/report
	CensysUrlReport string = "https://www.censys.io/api/v1/report/"

	// Censys supported standard user-agent header
	CensysHeaderUserAgent string = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36"

	// Censys supported standard accept header
	CensysHeaderAccept string = "application/json, */8"

	// define allowed index method for censys search api
	CensysIndexIPV4        string = "ipv4"
	CensysIndexCertificate string = "certificates"
	CensysIndexWebsites    string = "websites"
)

type CensysJson struct {
	Query   string   `json:"query"`
	Pages   int      `json:"pages,omitempty"`
	Fields  []string `json:"fields,omitempty"`
	Flatten bool     `json:"flatten,omitempty"`
}

type CensysSearchJson struct {
	Status  string `json:"status"`
	Results []struct {
		LocationCountry               string   `json:"location.country"`
		LocationRegisteredCountry     string   `json:"location.registered_country"`
		LocationLongitude             float64  `json:"location.longitude"`
		LocationCity                  string   `json:"location.city,omitempty"`
		IP                            string   `json:"ip"`
		LocationRegisteredCountryCode string   `json:"location.registered_country_code"`
		LocationCountryCode           string   `json:"location.country_code"`
		LocationLatitude              float64  `json:"location.latitude"`
		LocationProvince              string   `json:"location.province,omitempty"`
		LocationContinent             string   `json:"location.continent"`
		LocationPostalCode            string   `json:"location.postal_code,omitempty"`
		Protocols                     []string `json:"protocols"`
		LocationTimezone              string   `json:"location.timezone,omitempty"`
	} `json:"results"`
	Metadata struct {
		Count       int    `json:"count"`
		Query       string `json:"query"`
		BackendTime int    `json:"backend_time"`
		Page        int    `json:"page"`
		Pages       int    `json:"pages"`
	} `json:"metadata"`
}
