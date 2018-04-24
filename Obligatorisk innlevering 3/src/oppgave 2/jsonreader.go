package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
)

type Message1 struct {
	Ip string `json:"ip"`
}

type Message2 struct {
	Time string `json:"time"`
}

type Message3 struct {
	Country string `json:"country"`
	City string `json:"city"`
}

type Message4 struct {
	Card struct {
		Name string `json:"name"`
		Number int `json:"nationalPokedexNumber"`
	}
}

type Message5 struct {
	Text string `json:"text"`
}

func main() {
	http.HandleFunc("/", mainMsg)
	http.HandleFunc("/1", handler)
	http.HandleFunc("/2", handler2)
	http.HandleFunc("/3", handler3)
	http.HandleFunc("/4", handler4)
	http.HandleFunc("/5", handler5)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func mainMsg(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Please access one of the following pages, /1, /2, /3, /4, /5")
}

func handler(w http.ResponseWriter, r *http.Request) {
	res, _ := http.Get("https://api.ipify.org/?format=json")
	ip, _ := ioutil.ReadAll(res.Body)

	var msg Message1
	json.Unmarshal(ip, &msg)
	fmt.Fprintf(w, "Your IP is: %s!", msg.Ip)
}

func handler2(w http.ResponseWriter, r *http.Request) {
	res, _ := http.Get("http://date.jsontest.com/")
	ip, _ := ioutil.ReadAll(res.Body)

	var msg Message2
	json.Unmarshal(ip, &msg)
	fmt.Fprintf(w, "The Time: %s", msg.Time)
}

func handler3(w http.ResponseWriter, r *http.Request) {
	res, _ := http.Get("https://api.myjson.com/bins/1974az")
	ip, _ := ioutil.ReadAll(res.Body)
	var msg Message3
	json.Unmarshal(ip, &msg)
	fmt.Fprintf(w, "I'm from %s, in a small city called %s", msg.Country, msg.City)
}

func handler4(w http.ResponseWriter, r *http.Request) {
	res, _ := http.Get("https://api.pokemontcg.io/v1/cards/ex12-10")
	ip, _ := ioutil.ReadAll(res.Body)
	var msg Message4
	json.Unmarshal(ip, &msg)
	fmt.Fprintf(w, "My favorite pokemon is %s! It is number %d in the pokedex!\n", msg.Card.Name, msg.Card.Number)
}

func handler5(w http.ResponseWriter, r *http.Request) {
	res, _ := http.Get("https://api.myjson.com/bins/u8jez")
	ip, _ := ioutil.ReadAll(res.Body)
	var msg Message5
	json.Unmarshal(ip, &msg)
	fmt.Fprintf(w, "%s", msg.Text)
}
