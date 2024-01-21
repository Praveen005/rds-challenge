package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

)

const (
	getURL	  = "https://exam.ankush.wiki/assignments"
	baseURL   = "https://exam.ankush.wiki/data?part="
	cookieHdr = "connect.sid=s%3AKsYh6-5fnCkSPJzJx5sniFtHsHP8d.wiftarNatwREhke%2BIZc%2F2zU64%2FY1%2BJb76dK1dhbIHEk"
)

var ChainCode string

func getDataCount(url string)(int64){

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0
	}

	req.Header.Set("Cookie", cookieHdr)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil{
		return 0
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil{
		return 0
	}

	type getStruct struct {
		DataCount int64  `json:"numParts"`
		Message   string `json:"message"`
	}
	var val getStruct
	err = json.Unmarshal(data, &val)
	if err != nil{
		return 0
	}
	return val.DataCount
}

func downloadData(partNumber int, client *http.Client) ([]byte, error) {
	url := fmt.Sprintf("%s%d", baseURL, partNumber)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Cookie", cookieHdr)

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func getDataFunc() {
	client := &http.Client{
		Timeout: time.Second * 10, 
	}
	numberOfDataDownloads := getDataCount(getURL)
	for partNumber := 1; partNumber <= int(numberOfDataDownloads); partNumber++ {
		data, err := downloadData(partNumber, client)
		if err != nil {
			log.Printf("Error downloading part %d: %v", partNumber, err)
			continue
		}

		fmt.Printf("Downloaded data for part %d: %s\n", partNumber, string(data))
		JsonArray = append(JsonArray, string(data))

		time.Sleep(time.Second * 3)
	}

	ChainCode = ProcessJSONArray(JsonArray)
}
