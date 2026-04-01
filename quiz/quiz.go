package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Unable to parse CSV "+filePath, err)
	}
	return records
}

func main() {
	csvFile := flag.String("csv", "problems.csv", "CSV file for quiz")
	t := flag.Int("timer", 30, "Timer for quiz")
	s := flag.Bool("shuffle", false, "Shuffle the questions")
	flag.Parse()
	timer := time.NewTimer(time.Second * time.Duration(*t))
	records := readCsvFile(*csvFile)
	if *s == true {
		for i := range records {
			j := rand.Intn(i + 1)
			records[i], records[j] = records[j], records[i]
		}
	}
	var counter_correct int
	var counter_incorrect int

	for _, question := range records {
		fmt.Printf(question[0]+"=?   counter_correct=%d    counter_incorrect=%d\n", counter_correct, counter_incorrect)

		answerCh := make(chan string)

		go func() {
			var val string
			fmt.Scanln(&val)
			answerCh <- val
		}()

		select {
		case <-timer.C:
			fmt.Println("\nTimer's up!")
			fmt.Printf("Correct answers: %d\nIncorrect answers: %d\n", counter_correct, counter_incorrect)
			return
		case val := <-answerCh:
			if val == strings.TrimSpace(question[1]) {
				counter_correct++
				fmt.Println("Correct")
			} else {
				counter_incorrect++
				fmt.Println("Incorrect")
			}
		}

	}
	fmt.Printf("Correct answears: %d\n Incorrect answears: %d\n", counter_correct, counter_incorrect)
}
