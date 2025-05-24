package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// Flags
	fileName := flag.String("file", "problems.csv", "CSV file containing quiz questions")
	timeLimit := flag.Int("limit", 30, "Time limit for the quiz in seconds")
	flag.Parse()

	// Open CSV file
	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Read CSV records
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV data: %v\n", err)
		os.Exit(1)
	}

	type question struct {
		q string
		a string
	}
	var questions []question
	for _, record := range records {
		if len(record) == 2 {
			questions = append(questions, question{
				q: record[0],
				a: strings.TrimSpace(record[1]),
			})
		}
	}

	if len(questions) == 0 {
		fmt.Println("❌ No questions found in the CSV file.")
		return
	}

	// Wait for user to start quiz
	fmt.Printf("Press ENTER to start the quiz (you have %d seconds)...", *timeLimit)
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	// Setup timer
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	answerCh := make(chan string)
	exitCh := make(chan bool)

	// Input reader goroutine
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			ans, _ := reader.ReadString('\n')
			ans = strings.TrimSpace(ans)
			select {
			case answerCh <- ans:
			case <-exitCh:
				return
			}
		}
	}()

	// Quiz loop
	correct := 0
	total := len(questions)

	go func() {
		for i, q := range questions {
			fmt.Printf("Question #%d: %s = ", i+1, q.q)
			select {
			case userAns := <-answerCh:
				if userAns == q.a {
					correct++
				}
			case <-timer.C:
				fmt.Fprintln(os.Stderr, "\n⏰ Time is up!")
				close(exitCh)
				printResult(correct, total)
				return
			}
		}
		close(exitCh)
		printResult(correct, total)
	}()

	<-exitCh
}

func printResult(correct, total int) {
	// Force flush stdout by sleeping and using Fprintln
	fmt.Fprintf(os.Stderr, "✅ You scored %d out of %d.\n", correct, total)
	time.Sleep(500 * time.Millisecond) // Extra delay to ensure visibility
	os.Exit(0)
}
