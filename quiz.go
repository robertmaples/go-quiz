package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {
	filename := flag.String("file", "./problems.csv", "a csv file in the format of 'question,answer")
	sec := flag.Int("timer", 30, "time limit on the quiz in seconds")
	flag.Parse()

	f, err := os.Open(*filename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *filename))
	}

	reader := csv.NewReader(f)
	lines, err := reader.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	problems := parseLines(lines)
	correct := 0

	timer := time.NewTimer(time.Duration(*sec) * time.Second)
	go func() {
		<-timer.C
		fmt.Printf("\n\nTime has expired!\n\nYou got %d/%d correct!\n", correct, len(problems))
		os.Exit(0)
	}()

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)

		if answer == p.a {
			correct++
		}
	}
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}

	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
