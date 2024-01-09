package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	baseURL   = "https://exam.ankush.wiki/data?part="
	cookieHdr = "connect.sid=s%3AShnZDFr0rDdh-MPuhuym5ZwloAdckieu.PCQvfvexZw5CISOiluJdiPc%2BZNQ8UpQo%2F%2BEn%2Fe%2B7eI0"
)

var ChainCode string

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

	for partNumber := 1; partNumber <= 29; partNumber++ {
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
	fmt.Println("Mission completed successfully!")
}
