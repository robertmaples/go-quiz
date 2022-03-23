package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {
	filename := flag.String("file", "./problems.csv", "a csv file in the format of 'question,answer")
	timeLimit := flag.Int("limit", 30, "time limit on the quiz in seconds")
	shuffle := flag.Bool("shuffle", false, "set shuffle flag to randomize question order")
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

	if *shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(lines), func(i, j int) { lines[i], lines[j] = lines[j], lines[i] })
	}

	fmt.Print("Press the return key to begin the quiz.")
	fmt.Scanf("\n")
	fmt.Println()

	problems := parseLines(lines)
	correct := 0

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)

		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\n\nTime has expired!\n\nYou got %d/%d correct!\n", correct, len(problems))
			return
		case answer := <-answerCh:
			if strings.TrimSpace(answer) == p.a {
				correct++
			}
		}
	}

	fmt.Printf("\n\nYou got %d/%d correct!\n", correct, len(problems))
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

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
