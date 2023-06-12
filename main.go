package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey there, Welcome to our weather app\n")
	fmt.Fprintf(w, "To get status of weather in your city hit the route localhost:3000/weather/{citynam}")
}

func weather(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	cityName := params["city"]
	API_KEY := "9b7324d638sad8a20c080311b12c9das73adasa1ca7d0asasd6c37"
	get_Lat_Long := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&limit=1&appid=%s", cityName, API_KEY)
	resp, err := http.Get(get_Lat_Long)
	if err != nil {
		log.Fatal("Error in fetching lat and lon from Geocoding API")
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error in reading the response body")
	}
	var mp []map[string]interface{}
	if err := json.Unmarshal([]byte(bytes), &mp); err != nil {
		log.Fatal("Error while unmarshalling")
	}
	lat := mp[0]["lat"]
	lon := mp[0]["lon"]
	web_url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&appid=%s", lat, lon, API_KEY)
	response, err := http.Get(web_url)
	if err != nil {
		log.Fatal("Error while getting weather data")
	}
	val, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error in reading weather data")
	}
	fmt.Println(string(val))
	json.NewEncoder(w).Encode(string(val))
}

func main() {
	fmt.Println("Welcome to Weather tracker App")
	r := mux.NewRouter()
	r.HandleFunc("/", homePage)
	r.HandleFunc("/weather/{city}", weather)
	log.Fatal(http.ListenAndServe(":3000", r))
}
