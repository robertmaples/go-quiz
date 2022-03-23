package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	filename := flag.String("file", "./problems.csv", "specify a file path")
	flag.Parse()

	f, err := os.Open(*filename)
	if err != nil {
		log.Fatalln(err)
	}

	reader := csv.NewReader(f)

	var correct []bool
	for rec, recErr := reader.Read(); recErr != io.EOF; rec, recErr = reader.Read() {
		equation := rec[0]
		expected := rec[1]

		fmt.Printf("%s: ", equation)
		var actual string
		fmt.Scanf("%s\n", &actual)

		correct = append(correct, expected == actual)
	}

	var sum int
	for _, c := range correct {
		if c {
			sum += 1
		}
	}

	fmt.Printf("\nYou got %d/%d correct!\n", sum, len(correct))
}
