package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// flags,csv,os,channels and go rout for timer, time

//CSV data into a struct and usual data
type question struct {
	Quest string
	Anser int
}

const totalTime = 10

func main() {
	score := 0
	timer := time.NewTimer(time.Second * totalTime)
	defer timer.Stop()

	go func() {
		<-timer.C
		fmt.Println("Timer Alert")
		fmt.Print("Your score is: ")
		fmt.Print(score)
		fmt.Print("/13")
		fmt.Println()
		os.Exit(0)
	}()

	//Opens
	file, err := os.Open("problems.csv")
	if err != nil {

		log.Fatal(err)
	}
	//Closes
	defer file.Close()
	// Reads entire file
	csvReader, err := csv.NewReader(file).ReadAll()

	var questAns []question
	for _, csvReader := range csvReader {
		// look up str conv atoi
		ansInt, _ := strconv.Atoi(csvReader[1])
		data := question{
			Quest: csvReader[0],
			Anser: ansInt,
		}
		questAns = append(questAns, data)
	}

	for i := 0; i < len(questAns); i++ {
		var input int
		fmt.Println("What is the answer to " + questAns[i].Quest + "?")
		fmt.Scan(&input)
		if input == questAns[i].Anser {
			score++
		}

	}

}
