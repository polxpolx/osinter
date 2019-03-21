package havebeenpwned

import "time"

type HaveBeenPwnedJson struct {
	Name         string    `json:"Name"`
	Title        string    `json:"Title"`
	Domain       string    `json:"Domain"`
	BreachDate   string    `json:"BreachDate"`
	AddedDate    time.Time `json:"AddedDate"`
	ModifiedDate time.Time `json:"ModifiedDate"`
	PwnCount     int       `json:"PwnCount"`
	Description  string    `json:"Description"`
	LogoPath     string    `json:"LogoPath"`
	DataClasses  []string  `json:"DataClasses"`
	IsVerified   bool      `json:"IsVerified"`
	IsFabricated bool      `json:"IsFabricated"`
	IsSensitive  bool      `json:"IsSensitive"`
	IsRetired    bool      `json:"IsRetired"`
	IsSpamList   bool      `json:"IsSpamList"`
}
