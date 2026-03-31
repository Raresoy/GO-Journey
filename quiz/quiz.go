package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
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
	flag.Parse()
	records := readCsvFile(*csvFile)
	var counter_correct int
	var counter_incorrect int
	var val string

	for _, question := range records {
		fmt.Printf(question[0]+"=?   counter_correct=%d    counter_incorrect=%d\n", counter_correct, counter_incorrect)
		fmt.Scanln(&val)
		if val == strings.TrimSpace(question[1]) {
			counter_correct++
			fmt.Println("Correct")
		} else {
			fmt.Println("Incorrect")
			counter_incorrect++
		}
	}
	fmt.Printf("Correct answears: %d\n Incorrect answears: %d\n", counter_correct, counter_incorrect)
}
