package main

func getActiveCases() string {
	var url string = "https://api.covid19api.com/summary"
	var result string = consumeRestAPI(url)
	
	return result
}

func getAllCountries() string{
	var url string = "https://api.covid19api.com/countries"
	var result string = consumeRestAPI(url)

	return result
}

