package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"strconv"

)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func sumfromfile()  {
	lines, err := readLines("stdout.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	input1 := lines[0]
	input2 := lines[1]
	tall1,_ := strconv.Atoi(input1)
	tall2,_ := strconv.Atoi(input2)
	result := tall1 + tall2
	f, err := os.OpenFile("stdout.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := fmt.Fprintf(f,"\n%d\n", result); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
