package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, "http://") == false {
			aux := "http://"
			aux = aux + url
			url = aux
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			continue
		}
		_, error := io.Copy(os.Stdout, resp.Body)
		fmt.Println(resp.Status)
		resp.Body.Close()
		if error != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}

}
