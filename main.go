package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"strconv"
	//"encoding/json"
)

//Function to Consume REST API with URL input
func consumeRestAPI(url string) string{
	var input string = url
	var result string
	response, err := http.Get(input)
	if err != nil{
		log.Println(err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		result = string(data)
	}
	return result
}

func synchActiveCases(w http.ResponseWriter, r *http.Request){
	var activeCases string = ""
	activeCases = getActiveCases()
	createFile()
	writeFile(activeCases)
}

func getCases(w http.ResponseWriter, r *http.Request){
	queryParam := r.FormValue("country")
	number := 0
	response := ""

	//result := getActiveCases()

	if queryParam != "" {
		//var data1 = []activeCasesGlobal{}
		//err := json.Unmarshal([]byte(result), &data1)
		//	if err != nil {
		//		fmt.Println("error:", err)
		//}

		//fmt.Println("Test 1 : " , data1[0])
		//fmt.Println("Test 2 : " , data1)

		response = queryParam + " Active Cases : " + strconv.FormatInt(int64(number), 10)
	} else {
		//var data2 = map[string][]string{}
		//err := json.Unmarshal([]byte(result), &data2)
		//	if err != nil {
		//		fmt.Println("error:", err)
		//}
		response = "Total Active Cases : " + strconv.FormatInt(int64(number), 10)
	}
	
	w.Write([]byte(response))
}

func getDeaths(w http.ResponseWriter, r *http.Request){
	queryParam := r.FormValue("country")
	number := 0
	response := ""

	//result := getActiveCases()

	if queryParam != "" {
		//var data1 = []activeCasesGlobal{}
		//err := json.Unmarshal([]byte(result), &data1)
		//	if err != nil {
		//		fmt.Println("error:", err)
		//}

		//fmt.Println("Test 1 : " , data1[0])
		//fmt.Println("Test 2 : " , data1)

		response = queryParam + " Deaths Number : " + strconv.FormatInt(int64(number), 10)
	} else {
		//var data2 = map[string][]string{}
		//err := json.Unmarshal([]byte(result), &data2)
		//	if err != nil {
		//		fmt.Println("error:", err)
		//}
		response = "Total Deaths : " + strconv.FormatInt(int64(number), 10)
	}
	
	w.Write([]byte(response))
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/synchActiveCases", synchActiveCases)
	//http.HandleFunc("/deaths", getDeaths)
	//http.HandleFunc("/deathsTotal", getDeaths)
	http.HandleFunc("/cases", getCases)
	http.HandleFunc("/casesTotal", getCases)
	log.Fatal(http.ListenAndServe(":8081",nil))
}

func main(){
	handleRequests()
}