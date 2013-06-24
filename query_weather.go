package main

import (
	"encoding/json"
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
	requestUrl := fmt.Sprintf(FORECAST_API, apiKey, lat, lng)

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
	res := make(map[string][]map[string]map[string]map[string]interface{}, 0)
	json.Unmarshal(data, &res)

	lat, _ := res["results"][0]["geometry"]["location"]["lat"].(float64)
	lng, _ := res["results"][0]["geometry"]["location"]["lng"].(float64)

	return lat, lng
}

func extractWeatherInfo(data []byte) string {
	var res ForecastResults
	json.Unmarshal(data, &res)

	return res.Daily.Summary
}

func main() {
	lat, lng := getGeoCode("Berlin, Germany")
	fmt.Println(getWeatherInfo(lat, lng))
}
