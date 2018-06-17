package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// Response Object using as the core struct to map the json response too
type Response struct {
	Start *time.Time `json:"start,omitempty"`
	End   *time.Time `json:"end,omitempty"`
	Data  Data       `json:"data"`
}

// Data structure, intended to be embdeed to the response Object
type Data struct {
	Standard float32 `json:"standard"`
	Reduced  float32 `json:"reduced"`
}

var date string
var vatType string

func main() {
	arguments := os.Args[1:]

	// ensure only valid parameters are parsed
	_, err := validateParams(arguments)
	if err != nil {
		log.Fatal(err)
	}

	vatType = arguments[0]
	date = arguments[1]
	currentDate, _ := time.Parse(time.RFC3339, date)

	// calls the api response and returns a struct of type Response
	response, err := getJSONResponse("http://localhost:8080/api/1.0/views/uk-vat-rates")
	if err != nil {
		log.Fatal(err)
	}

	// takes []Response and loops through to find any matching
	for _, response := range response {
		if response.End == nil {
			var t = time.Now()
			t.Format(time.RFC3339)
			response.End = &t
		}

		if currentDate.Unix() > response.Start.Unix() && currentDate.Unix() < response.End.Unix() {
			if vatType == "standard" {
				fmt.Println("The standard rate of VAT for", currentDate, "is", response.Data.Standard*100, "%")
			} else {
				fmt.Println("The reduced rate of VAT for", currentDate, "is", response.Data.Reduced*100, "%")
			}
		}
	}
}

// This function is designed to valid the Arguments slice, ensuring the user cannot
// enter invalid arguments. Also delivering an error message if needed.
func validateParams(arguments []string) (bool, error) {
	// checks to ensure two params are given
	if len(arguments) != 2 {
		return false, errors.New("There needs to be two arguments for this fucntion to work")
	}
	// checks to ensure only standard or reduced are parsed
	vat := strings.Trim(arguments[0], "\n")
	if vat != "standard" && vat != "reduced" {
		return false, errors.New("The First parameter needs to be either standard or reduced")
	}

	_, err := time.Parse(time.RFC3339, arguments[1])
	if err != nil {
		return false, errors.New("There was a problem with the date you provided")
	}

	return true, nil
}

// This function is intended to use  go's Http client to return the response from the server
// Upon a successful call. The response is bound to a slice of Response and returned.
func getJSONResponse(url string) ([]Response, error) {
	var response = []Response{}
	var httpClient = &http.Client{
		Timeout: time.Second * 10,
	}

	res, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response, err
}
