package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
	"runtime"
	"os/exec"
	"html/template"
)

//Definerer structs for json-innlesing
type WeatherData struct {
	Weather struct {
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

//Lytter til port og åpner nettleser
func main() {
	nettleser("http://localhost:8080/")
	http.HandleFunc("/", mainMsg)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//Åpner nettleser
func nettleser(url string) bool {
	var args []string
	switch runtime.GOOS {
	case "windows":
		args = []string{"cmd", "/c", "start"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}

//Henter API fra Kristiansand ved cityID 6453405
func mainMsg(w http.ResponseWriter, r *http.Request) {
	res, _ := http.Get("http://api.openweathermap.org/data/2.5/weather?id=6453405&APPID=2824f4858c14c0799bb09f88fee6a7f0")
	weather, _ := ioutil.ReadAll(res.Body)
	var msg WeatherData
	json.Unmarshal(weather, &msg)


// Gjør om temperaturen til Celsius
	msg.Main.Temp = msg.Main.Temp - 273.15

//if-statements for å bestemme bekledning, sett opp mot temperaturen i Kristiansand
	if (msg.Main.Temp > 18){

		msg.Weather.Description = fmt.Sprintf("Temperaturen i dag er %2.1f grader! Finn frem shortsen, det er sommer i %s!" +"\n",msg.Main.Temp, msg.Name)
		if (msg.Main.Humidity > 80){
			msg.Weather.Description += fmt.Sprintf("Høy luftfuktighet i dag, hele %d!",msg.Main.Humidity)
		}
		if (msg.Main.Humidity < 30){
			msg.Weather.Description += fmt.Sprintf("Lav luftfuktighet i dag, så lavt som %d!",msg.Main.Humidity)
		}
		if (msg.Main.Humidity < 80 && msg.Main.Humidity > 30){
			msg.Weather.Description += fmt.Sprintf("Helt moderat luftfuktighet i dag, %d.",msg.Main.Humidity)
		}
	}
	if (msg.Main.Temp < 18 && msg.Main.Temp > 10){

		msg.Weather.Description = fmt.Sprintf("Temperaturen i dag er %2.1f grader! Det nærmer seg sommervær i %s!" +"\n",msg.Main.Temp, msg.Name)
		if (msg.Main.Humidity > 80){
			msg.Weather.Description += fmt.Sprintf("Høy luftfuktighet i dag, hele %d!",msg.Main.Humidity)
		}
		if (msg.Main.Humidity < 30){
			msg.Weather.Description += fmt.Sprintf("Lav luftfuktighet i dag, så lavt som %d!",msg.Main.Humidity)
		}
		if (msg.Main.Humidity < 80 && msg.Main.Humidity > 30){
			msg.Weather.Description += fmt.Sprintf("Helt moderat luftfuktighet i dag, %d.",msg.Main.Humidity)
		}
	}
	if (msg.Main.Temp < 10) {

		msg.Weather.Description = fmt.Sprintf("Langbuksevær i dag! Temperaturen er %2.1f grader! Krysser fingrene for bedre vær i %s." +"\n",msg.Main.Temp, msg.Name)
		if (msg.Main.Humidity > 80){
			msg.Weather.Description += fmt.Sprintf("Høy luftfuktighet i dag, hele %d! Kanskje det er lurt å ta med paraply om du skal ut?",msg.Main.Humidity)
		}
		if (msg.Main.Humidity < 30){
			msg.Weather.Description += fmt.Sprintf("Lav luftfuktighet i dag, så lavt som %d!",msg.Main.Humidity)
		}
		if (msg.Main.Humidity < 80 && msg.Main.Humidity > 30){
			msg.Weather.Description += fmt.Sprintf("Helt moderat luftfuktighet i dag, %d.",msg.Main.Humidity)
		}
	}
	//Mater data til Warm.html
	temp, err := template.ParseFiles("Warm.html")
	if err != nil {
		log.Print(err)
	}

	err = temp.Execute(w, msg)
	if err != nil {
		log.Fatal(err)
	}
}
