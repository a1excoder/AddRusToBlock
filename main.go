package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const url = "https://uablacklist.net/all.json"

type DataBlock struct {
	Alias string   `json:"alias"`
	Term  string   `json:"term"`
	Urls  []string `json:"urls"`
	Ips   []string `json:"ips"`
}

type DataBlockMain map[string]DataBlock

func (data *DataBlock) ViewDataBlock() {
	fmt.Printf("\talias: %s | term: %s\n", data.Alias, data.Term)
	fmt.Println("\turls:")
	for i, url := range data.Urls {
		fmt.Printf("\t\t%d) %s\n", i, url)
	}

	fmt.Println("\n\tips:")
	for i, url := range data.Ips {
		fmt.Printf("\t\t%d) %s\n", i, url)
	}

	fmt.Println("")
}

func main() {
	response, err := http.Get(url)
	if err != nil {
		log.Fatalln(err.Error())
	}

	bodyData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err.Error())
	}

	blockList := DataBlockMain{}
	if err = json.Unmarshal(bodyData, &blockList); err != nil {
		log.Fatalln(err.Error())
	}

	for name, block := range blockList {
		fmt.Printf("name: %s\n", name)
		block.ViewDataBlock()
	}
}
