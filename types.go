package main

import (
	"time"
)

type Response struct {
	Message   string      `json:"Message"`
	Global    Global      `json:"Global"`
	Countries []Countries `json:"Countries"`
}
type Global struct {
	NewConfirmed   int `json:"NewConfirmed"`
	TotalConfirmed int `json:"TotalConfirmed"`
	NewDeaths      int `json:"NewDeaths"`
	TotalDeaths    int `json:"TotalDeaths"`
	NewRecovered   int `json:"NewRecovered"`
	TotalRecovered int `json:"TotalRecovered"`
}
type Countries struct {
	Country        string    `json:"Country"`
	CountryCode    string    `json:"CountryCode"`
	Slug           string    `json:"Slug"`
	NewConfirmed   int       `json:"NewConfirmed"`
	TotalConfirmed int       `json:"TotalConfirmed"`
	NewDeaths      int       `json:"NewDeaths"`
	TotalDeaths    int       `json:"TotalDeaths"`
	NewRecovered   int       `json:"NewRecovered"`
	TotalRecovered int       `json:"TotalRecovered"`
	Date           time.Time `json:"Date"`
}
