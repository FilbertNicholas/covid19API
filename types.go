package main

type activeCaseGlobal struct {
	Global  struct {
		NewConfirmed   int `json:"NewConfirmed"`
		TotalConfirmed int `json:"TotalConfirmed"`
		NewDeaths      int `json:"NewDeaths"`
		TotalDeaths    int `json:"TotalDeaths"`
		NewRecovered   int `json:"NewRecovered"`
		TotalRecovered int `json:"TotalRecovered"`
	} `json:"Global"`
}

type activeCasePerCountry struct {
	Countries []struct {
		Country        string    `json:"Country"`
		CountryCode    string    `json:"CountryCode"`
		Slug           string    `json:"Slug"`
		NewConfirmed   int       `json:"NewConfirmed"`
		TotalConfirmed int       `json:"TotalConfirmed"`
		NewDeaths      int       `json:"NewDeaths"`
		TotalDeaths    int       `json:"TotalDeaths"`
		NewRecovered   int       `json:"NewRecovered"`
		TotalRecovered int       `json:"TotalRecovered"`
	}
}

//type Death struct {
//	Title string "title"
//	Desc string "desc"
//	Message string "message"
//}

type activeCasesGlobal []activeCaseGlobal
type activeCasesPerCountry []activeCasePerCountry
//type Deaths []Death