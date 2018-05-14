package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	dumpRaw = false
	zip     = "4360,no"
	api     = "b36acf3e5ce6ef0067244e002ca8c4a0"
)

var (
	weatherKeys = map[string]bool{"main": false, "wind": false, "coord": false, "weather": true, "sys": false, "clouds": false}
)

func main() {

	urlString := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?zip=%s&APPID=%s", zip, api)
	u, err := url.Parse(urlString)
	res, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}

	jsonBlob, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data map[string]interface{}

	if dumpRaw {
		fmt.Printf("blob = %s\n\n", jsonBlob)
	}
	err = json.Unmarshal(jsonBlob, &data)
	if err != nil {
		fmt.Println("error:", err)
	}

	if dumpRaw {
		fmt.Printf("%+v", data)
	}
	for k, v := range data {
		val, isAnArray := isKey(k)
		if val {
			dumpMap(k, v, isAnArray)
		} else {
		}
	}
}

func dumpMap(k string, v interface{}, isArray bool) {

	fmt.Printf("%s:\n", k)
	if isArray {
		for i := 0; i < len(v.([]interface{})); i++ {
			nmap := v.([]interface{})[i].(map[string]interface{})
			for kk, vv := range nmap {
				fmt.Printf("\tthe %s is %v\n", kk, vv)
			}
		}
	} else {
		nmap := v.(map[string]interface{})
		for kk, vv := range nmap {
			if isTempVal(kk) {
				farenTemp := faren(vv.(float64))
				fmt.Printf("\tthe %s is %f\n", kk, farenTemp)
			} else if isSunVal(kk) {
				sunTime := time.Unix((int64(vv.(float64))), 0)
				fmt.Printf("\tthe %s at %v\n", kk, sunTime)
			} else {
				fmt.Printf("\tthe %s is %v\n", kk, vv)
			}
		}
	}
}

func isKey(k string) (ok bool, isArray bool) {
	isArray, ok = weatherKeys[k]
	return ok, isArray
}

func faren(kelvin float64) float64 {
	return (9.0/5.0*(kelvin-273.0) + 32.0)
}

func isTempVal(k string) bool {
	return strings.Contains(k, "temp")
}

func isSunVal(k string) bool {
	return strings.Contains(k, "sun")
}