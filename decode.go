package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"strconv"
)

type DataStruct struct {
	Data []string `json:"data"`
}

var JsonArray []string

var morse = map[string]string{
	"-----": "0",
	".----": "1",
	"..---": "2",
	"...--": "3",
	"....-": "4",
	".....": "5",
	"-....": "6",
	"--...": "7",
	"---..": "8",
	"----.": "9",
	".-": "A",
	"-...": "B",
	"-.-.": "C",
	"-..": "D",
	".": "E",
	"..-.": "F",
	"--.": "G",
	"....": "H",
	"..": "I",
	".---": "J",
	"-.-": "K",
	".-..": "L",
	"--": "M",
	"-.": "N",
	"---": "O",
	".--.": "P",
	"--.-": "Q",
	".-.": "R",
	"...": "S",
	"-": "T",
	"..-": "U",
	"...-": "V",
	".--": "W",
	"-..-": "X",
	"-.--": "Y",
	"--..": "Z",
}

func breakString(input string, maxLength int) []string {
	var result []string

	for i := 0; i < len(input); i += maxLength {
		end := i + maxLength
		if end > len(input) {
			end = len(input)
		}
		result = append(result, input[i:end])
	}

	return result
}

func decodeCode(jsonData string) (string, string){

	// Unmarshal JSON data
	var dataStruct DataStruct
	err := json.Unmarshal([]byte(jsonData), &dataStruct)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return "", ""
	}

	// Extract "data" field
	data := dataStruct.Data
	dataString :=strings.Join(data, "")
	arrowIndex := strings.Index(dataString, "➡➡➡")
	leftSide := strings.TrimSpace(dataString[:arrowIndex])
	rightSide := strings.TrimSpace(dataString[arrowIndex+len("➡➡➡"):])

	var leftMcodes []string
	var rightMcodes []string

	leftMcodes = breakString(leftSide, 5)
	rightMcodes = breakString(rightSide, 5)


	var leftDecode []string
	var rightDecode string

	for _, val := range leftMcodes{
		leftDecode = append(leftDecode, morse[val])
	}
	// fmt.Println("")
	for _, val := range rightMcodes{
		rightDecode = morse[val]
	}

	s1 := strings.Join(leftDecode, "")


	return s1, rightDecode
}

func ProcessJSONArray(jsonArray []string){

	var entries []struct{
		key string
		value string
	}

	

	for _, val := range jsonArray{
		lValue, rValue := decodeCode(val)

		entries = append(entries, struct{key string; value string} {lValue, rValue})
	}

	sort.Slice(entries, func(i, j int) bool {
		key1, _ := strconv.Atoi(entries[i].key)
		key2, _ := strconv.Atoi(entries[j].key)
		return key1 < key2
	})

	var result string =""
	for _, value := range entries{
		fmt.Println(value.key, " - ", value.value)
		result = result + value.value
	}

	fmt.Println(result);
}



var _ = morse
var _ = JsonArray
var _ = breakString
var _ = decodeCode