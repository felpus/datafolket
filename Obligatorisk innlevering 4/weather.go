package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
)

//Definerer structs for json-innlesing
type WeatherData struct {
	Weather [] struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`


	Main struct {
		Temp     float32 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`

	} `json:"main"`

	Sys struct {
		ID      int     `json:"id"`
		Country string  `json:"country"`
	} `json:"sys"`
	Name string `json:"name"`
}

//Lytter til port
func main() {
	http.HandleFunc("/", mainMsg)
	log.Fatal(http.ListenAndServe(":8080", nil))
}


//Henter API
func mainMsg(w http.ResponseWriter, r *http.Request) {
	res, _ := http.Get("http://api.openweathermap.org/data/2.5/weather?id=6453405&APPID=2824f4858c14c0799bb09f88fee6a7f0")
	weather, _ := ioutil.ReadAll(res.Body)
	var msg WeatherData
	json.Unmarshal(weather, &msg)


// Gjør om temperaturen til Celsius
	msg.Main.Temp = msg.Main.Temp - 273.15

//if-statements for å bestemme bekledning, sett opp mot temperaturen i Kristiansand
	if (msg.Main.Temp > 18){

		fmt.Fprintf(w, "Temperaturen i dag er %f! Finn frem shortsen, det er sommer i %s!",msg.Main.Temp, msg.Name)
		if (msg.Main.Humidity > 80){
			fmt.Fprintf(w, "Høy luftfuktighet i dag, hele %d!",msg.Main.Humidity)
		}
		if (msg.Main.Humidity < 30){
			fmt.Fprintf(w, "Lav luftfuktighet i dag, så lavt som %d!",msg.Main.Humidity)
		}
		if (msg.Main.Humidity < 80 && msg.Main.Humidity > 30){
			fmt.Fprintf(w, "Helt moderat luftfuktighet i dag, %d.",msg.Main.Humidity)
		}
	}
	if (msg.Main.Temp < 18 && msg.Main.Temp > 10){
		fmt.Fprintf(w, "Temperaturen i dag er %f! Det nærmer seg sommervær i %s!" +"\n",msg.Main.Temp, msg.Name)
		if (msg.Main.Humidity > 80){
			fmt.Fprintf(w, "Høy luftfuktighet i dag, hele %d!",msg.Main.Humidity)
		}
		if (msg.Main.Humidity < 30){
			fmt.Fprintf(w, "Lav luftfuktighet i dag, så lavt som %d!",msg.Main.Humidity)
		}
		if (msg.Main.Humidity < 80 && msg.Main.Humidity > 30){
			fmt.Fprintf(w, "Helt moderat luftfuktighet i dag, %d.",msg.Main.Humidity)
		}
	}
	if (msg.Main.Temp < 10) {
		fmt.Fprintf(w, "Langbuksevær i dag! Temperaturen er %f! Krysser fingrene for bedre vær i %s.",msg.Main.Temp, msg.Name)
		if (msg.Main.Humidity > 80){
			fmt.Fprintf(w, "Høy luftfuktighet i dag, hele %d!",msg.Main.Humidity)
		}
		if (msg.Main.Humidity < 30){
			fmt.Fprintf(w, "Lav luftfuktighet i dag, så lavt som %d!",msg.Main.Humidity)
		}
		if (msg.Main.Humidity < 80 && msg.Main.Humidity > 30){
			fmt.Fprintf(w, "Helt moderat luftfuktighet i dag, %d.",msg.Main.Humidity)
		}
	}

}
