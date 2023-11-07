package main

import (
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

	ApiKey := os.Getenv("API_KEY")

	response, err := http.Get("http://datapoint.metoffice.gov.uk/public/data/val/wxfcs/all/xml/3840?res=3hourly&key=" + ApiKey)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sbE := string(body)
	log.Printf(sbE)
}
