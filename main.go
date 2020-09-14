package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

//Function to Consume REST API with URL input
func consumeRestAPI(url string) string {
	var input string = url
	var result string
	response, err := http.Get(input)
	if err != nil {
		log.Println(err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		result = string(data)
	}
	return result
}

func synchActiveCases(w http.ResponseWriter, r *http.Request) {
	var activeCases string = ""
	activeCases = getActiveCases()
	createFile()
	writeFile(activeCases)

	w.Write([]byte("Done"))
}

func getCases(w http.ResponseWriter, r *http.Request) {
	queryParam := r.FormValue("country")
	number := 0
	response := ""

	result := getActiveCases()

	if queryParam != "" {
		var data Response
		err := json.Unmarshal([]byte(result), &data)
		if err != nil {
			fmt.Println("error:", err)
		}

		for _, i := range data.Countries {
			if i.CountryCode == queryParam {
				number = i.TotalConfirmed - i.TotalDeaths - i.TotalRecovered
			}
		}

		response = queryParam + " Active Cases : " + strconv.FormatInt(int64(number), 10)

	} else {
		var data Response
		err := json.Unmarshal([]byte(result), &data)
		if err != nil {
			fmt.Println("error:", err)
		}

		number = data.Global.TotalConfirmed - data.Global.TotalDeaths - data.Global.TotalRecovered
		response = "Total Active Cases : " + strconv.FormatInt(int64(number), 10)
	}

	w.Write([]byte(response))
}

func getDeaths(w http.ResponseWriter, r *http.Request) {
	queryParam := r.FormValue("country")
	number := 0
	response := ""

	result := getActiveCases()

	if queryParam != "" {
		var data Response
		err := json.Unmarshal([]byte(result), &data)
		if err != nil {
			fmt.Println("error:", err)
		}

		for _, i := range data.Countries {
			if i.CountryCode == queryParam {
				number = i.TotalDeaths
			}
		}

		response = queryParam + " Deaths : " + strconv.FormatInt(int64(number), 10)

	} else {
		var data Response
		err := json.Unmarshal([]byte(result), &data)
		if err != nil {
			fmt.Println("error:", err)
		}

		number = data.Global.TotalDeaths
		response = "Total Deaths : " + strconv.FormatInt(int64(number), 10)
	}

	w.Write([]byte(response))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/synchActiveCases", synchActiveCases)
	http.HandleFunc("/deaths", getDeaths)
	http.HandleFunc("/deathsTotal", getDeaths)
	http.HandleFunc("/cases", getCases)
	http.HandleFunc("/casesTotal", getCases)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
