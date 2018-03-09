package main

import (
	"fmt"
	"time"
	"os"
	"os/signal"
)

func main() {
	d := make(chan os.Signal, 1)
	signal.Notify(d, os.Interrupt)
	go func() {
		<-d
		fmt.Println("Feilmelding")
		os.Exit(1)
	}()
	c := make(chan int)
	go readInput(c)
	go addUp(c)
	time.Sleep(10 * 1e9)
}

func addUp(c chan int) {
	n1, n2 := <-c, <-c
	res := n1 + n2
	c <- res
}

func readInput(c chan int) {
	var n1 int
	var n2 int
	fmt.Println("Skriv verdi: ")
	fmt.Scan(&n1)
	fmt.Println("Skriv verdi: ")
	fmt.Scan(&n2)
	c <- n1
	c <- n2
	res := <-c
	fmt.Println("Resultat: ", res)
}

