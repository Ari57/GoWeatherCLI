package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ResponseResult struct {
	Geometry struct {
		Location struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
	} `json:"geometry"`
}

type Response struct {
	Results []ResponseResult `json:"results"`
}

func CallGeocodeApi(GeocodeKey string, location string) []byte {
	fmt.Println(GeocodeKey) // TODO remove
	HttpLink := "https://maps.googleapis.com/maps/api/geocode/json?key=" + GeocodeKey + "&address=" + location
	fmt.Println(HttpLink)

	resp, err := http.Get(HttpLink)
	ErrorChecker(err)

	responseData, err := io.ReadAll(resp.Body)
	ErrorChecker(err)

	return responseData
}

func FormatResult(responseData []byte) (Response, error) {
	var responseObject Response
	err := json.Unmarshal(responseData, &responseObject)

	return responseObject, err
}

func ReturnResult(responseObject Response) (string, string, error) {
	if len(responseObject.Results) > 0 {
		lat := responseObject.Results[0].Geometry.Location.Lat
		lng := responseObject.Results[0].Geometry.Location.Lng

		return fmt.Sprint(lat), fmt.Sprint(lng), nil
	}

	return "0", "0", fmt.Errorf("No results found")

}

func ErrorChecker(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
