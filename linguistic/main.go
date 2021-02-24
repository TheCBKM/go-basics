package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func getLanguages(url string) {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error languages not found")
	}
	var r map[string]float64
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		log.Fatalln(err)
	}
	sum := 0.0
	for key := range r {
		sum += r[key]
	}

	for key := range r {

		fmt.Printf("\n%v = %.2f %% \n", key, (r[key] / sum * 100))
	}
}

func main() {

	url := "https://github.com/TheCBKM/go-basics"

	strs := strings.Split(url, "/")
	if len(strs) != 5 {
		log.Fatal("Invalid URL")
	}

	resp, err := http.Get("https://api.github.com/repos/" + strs[len(strs)-2] + "/" + strs[len(strs)-1])
	if err != nil {
		log.Fatal("Error repo not found ")
	}
	var r map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		log.Fatalln(err)
	}

	url = fmt.Sprintf("%v", r["languages_url"])

	getLanguages(string(url))
}
