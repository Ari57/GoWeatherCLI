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

func main() {
	err := godotenv.Load("credential.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	GeocodeKey := os.Getenv("GEOCODE_KEY")
	ApiKey := os.Getenv("API_KEY")

	location := ""

	fmt.Print("Enter a location: ")
	fmt.Scan(&location)
	fmt.Println("Provided location is", location)

	lat, lng, err := Geocode(GeocodeKey, location)
	ErrorCheck(err)

	HttpLink := "https://api.openweathermap.org/data/3.0/onecall?lat=" + lat + "&lon=" + lng + "&appid=" + ApiKey + "&units=metric"

	resp, err := http.Get(HttpLink)
	ErrorChecker(err)

	responseData, err := io.ReadAll(resp.Body)
	ErrorChecker(err)

	type CurrentWeather struct {
		Temp float64 `json:"temp"`
	}

	type Response struct {
		Current CurrentWeather `json:"current"`
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	temperature := fmt.Sprint(responseObject.Current.Temp)
	fmt.Println("The current temperature is " + temperature + " Celsius")
}

func ErrorCheck(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
