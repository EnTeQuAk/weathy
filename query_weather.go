package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	GEOCODE_BASE_URL = "http://maps.googleapis.com/maps/api/geocode/json"
	FORECAST_API     = "https://api.forecast.io/forecast/%s/%f,%f?units=si"
)

type ForecastResults struct {
	Latitude  float64   `json:latitude`
	Longitude float64   `json:longitude`
	Timezone  string    `json:string`
	Daily     DailyType `json:daily`
}

type DailyType struct {
	Summary string `json:summary`
	Icon    string `json:icon`
}

func getGeoCode(city string) (float64, float64) {
	values := make(url.Values)
	values.Set("address", city)
	values.Set("sensor", "false")

	requestUrl := GEOCODE_BASE_URL + "?" + values.Encode()

	response, err := http.Get(requestUrl)

	if err != nil {
		log.Panic(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Panic(err)
	}

	lat, lng := extractLatLngFromResponse(body)

	return lat, lng
}

func getWeatherInfo(lat float64, lng float64) string {
	apiKey := os.Getenv("FORECAST_API_KEY")
	if apiKey == "" {
		log.Fatal("Please define FORECAST_API_KEY")
		os.Exit(1)
	}
	requestUrl := fmt.Sprintf(FORECAST_API, apiKey, lat, lng)

	log.Print(fmt.Sprintf("Fetch %s", requestUrl))

	response, err := http.Get(requestUrl)

	if err != nil {
		log.Panic(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Panic(err)
	}

	info := extractWeatherInfo(body)

	return info
}

func extractLatLngFromResponse(data []byte) (float64, float64) {
	type Location struct {
		Lat float64
		Lng float64
	}

	type Response struct {
		Results []struct {
			Geometry struct {
				Location Location
			}
		}
	}

	var resp Response
	json.Unmarshal(data, &resp)
	lat := resp.Results[0].Geometry.Location.Lat
	lng := resp.Results[0].Geometry.Location.Lng

	return lat, lng
}

func extractWeatherInfo(data []byte) string {
	var res ForecastResults
	json.Unmarshal(data, &res)

	return res.Daily.Summary
}

func main() {
	var city string

	flag.StringVar(&city, "city", "Berlin, Germany", "City for Forecast")
	flag.Parse()

	lat, lng := getGeoCode(city)
	fmt.Println(getWeatherInfo(lat, lng))
}
