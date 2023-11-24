package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type CurrentWeather struct {
	Temp float64 `json:"temp"`
}

type Result struct {
	Current CurrentWeather `json:"current"`
}

func main() {
	GeocodeKey, ApiKey := GetKeys()
	responseData := CallWeatherApi(GeocodeKey, ApiKey)
	FormatReturnResult(responseData)

}

func GetKeys() (string, string) {
	err := godotenv.Load("credential.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	GeocodeKey := os.Getenv("GEOCODE_KEY")
	ApiKey := os.Getenv("API_KEY")

	return GeocodeKey, ApiKey
}

func CallWeatherApi(GeocodeKey string, ApiKey string) []byte {
	location := ""

	fmt.Print("Enter a location: ")
	fmt.Scan(&location)
	fmt.Println("Provided location is", location)

	ResponseData := CallGeocodeApi("fbwfbwfbfuobwf", location)
	ResponseObject, err := FormatResult(ResponseData)
	ErrorCheck(err)
	lat, lng, err := ReturnResult(ResponseObject)
	ErrorCheck(err)

	HttpLink := "https://api.openweathermap.org/data/3.0/onecall?lat=" + lat + "&lon=" + lng + "&appid=" + ApiKey + "&units=metric"

	resp, err := http.Get(HttpLink)
	ErrorChecker(err)

	responseData, err := io.ReadAll(resp.Body)
	ErrorChecker(err)

	return responseData
}

func FormatReturnResult(responseData []byte) {
	var responseObject Result
	json.Unmarshal(responseData, &responseObject)
	temperature := fmt.Sprint(responseObject.Current.Temp)
	fmt.Println("The current temperature is " + temperature + " Celsius")

}

func ErrorCheck(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
