package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main() {
	file := "fisier.txt"
	f, err := os.Open(file)
	f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "eroare: %v\n", err)
	} else {
		input := bufio.NewScanner(f)
		for input.Scan() {
			resp, err := http.Get(input.Text())
			defer resp.Body.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "eroare: %v\n", err)
				continue
			}
			if resp.Status == "200 OK" {
				fmt.Println(input.Text())
			}
		}
	}
}
