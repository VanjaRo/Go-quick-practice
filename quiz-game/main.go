package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv",
		"a csv file wtih problems in the format of 'question, answer'")
	timeLimit := flag.Int("limit", 30, "limit for a question to answer in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("failed to open the %s file \n", *csvFilename))
		os.Exit(2)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("failed to parse the %s file \n", *csvFilename))
		os.Exit(2)
	}
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	score := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("\nYour time is out!")
			fmt.Printf("Your scored %d out of %d.\n", score, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.a {
				score++
			}
			timer.Reset(time.Duration(*timeLimit) * time.Second)
		}
	}
	fmt.Printf("Your scored %d out of %d.\n", score, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	log.Print(msg)
	os.Exit(2)
}
