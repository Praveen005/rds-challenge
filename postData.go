package main

import(
	"fmt"
	"net/http"
	"strings"
	"io"
)

const(
	submitURL = "https://exam.ankush.wiki/answers"
)


func SubmitChaincode() {

	dataToBeSent := fmt.Sprintf(`{ "chaincode": "%s" }`, ChainCode)

    // Making the POST request
    response, err := http.Post(submitURL, "application/json", strings.NewReader(dataToBeSent))
    if err != nil {
		fmt.Println(err.Error())

    }
    defer response.Body.Close()

    // Read the response from server
    responseBody, err := io.ReadAll(response.Body)
    if err != nil {
       fmt.Println(err)
    }

	//Printing the response
    fmt.Println("Response from server:\n", string(responseBody))

}