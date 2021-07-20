package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	filename := flag.String("f", "problems.csv", "The csv filename which contains the problem set")
	timeLimit := flag.Int("time", 30, "The amount of time to complete the quiz")

	flag.Parse()

	file, err := os.Open(*filename)

	if err != nil {
		log.Fatal(err)
	}

	csvLines, err := csv.NewReader(file).ReadAll()

	if err != nil {
		fmt.Println(err)
	}

	correctCount := 0
	answeredQuestions := 0
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("You have x amount of time complete this quiz, are you ready? Press any key to continue")
	reader.ReadString('\n')

	questionCount := len(csvLines)
	timer := time.AfterFunc(time.Duration(*timeLimit)*time.Second, func() {
		if answeredQuestions == questionCount {
			return
		} else {
			fmt.Printf("Times, up! You got %d/%d questions correct.", correctCount, questionCount)
			os.Exit(0)
		}
	})

	for _, line := range csvLines {
		answeredQuestions += 1
		fmt.Println("What's the answer to ", line[0])
		text, _ := reader.ReadString('\n')

		if line[1] == strings.TrimRight(text, "\r\n") {
			fmt.Println("Correct")
			correctCount += 1
		} else {
			fmt.Println("Incorrect")

		}

	}
	timer.Stop()

	fmt.Printf("You got %d/%d questions correct.", correctCount, questionCount)
}
