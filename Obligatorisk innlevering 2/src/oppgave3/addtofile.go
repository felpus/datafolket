package main

import (
	"os"
	"log"
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"os/signal"
)

func addtofile() {
	var n1 int
	var n2 int
	fmt.Println("Skriv verdi: ")
	fmt.Scan(&n1)
	fmt.Println("Skriv verdi: ")
	fmt.Scan(&n2)
	file, err := os.Create("stdout.txt")
	if err != nil {
		log.Fatal("Kan ikke opprette fil", err)
	}
	defer file.Close()
	f, err := os.OpenFile("stdout.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := fmt.Fprintf(f, "%d\n%d", n1, n2); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func readResult(path string) {
	data, err := ioutil.ReadFile(path)
	checkErr(err)
	tempData := string(data)
	stringData := strings.Split(tempData, "\n")
	temp := stringData[len(stringData)-2]
	resultat, err := strconv.Atoi(temp)
	checkErr(err)
	fmt.Println("Resultat: ", resultat)
}

func main() {
	d := make(chan os.Signal, 1)
	signal.Notify(d, os.Interrupt)
	go func() {
		<-d
		fmt.Println("Feilmelding")
		os.Exit(1)
	}()
	addtofile()
	sumfromfile()
	readResult("stdout.txt")
}