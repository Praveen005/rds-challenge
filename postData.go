package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const(
	submitURL = "https://exam.ankush.wiki/answers"
)


func SubmitChaincode(url string, ChainCode string)([]byte, error) {

    
    type respToSubmit struct{
      Chaincode string    `json:"chaincode"`
    }

    var val respToSubmit
    val.Chaincode = ChainCode
    serializedVal, err := json.Marshal(val)
    if err != nil{
      return nil, err
    }

    req, err := http.NewRequest("POST", url, strings.NewReader(string(serializedVal)))
    req.Header.Set("Content-Type", "application/json")
    if err != nil{
      return nil, err
    }
    req.Header.Set("Cookie", cookieHdr)
    client := &http.Client{}
    response, err := client.Do(req)

    if err != nil {
      return nil, err
    }
    defer response.Body.Close()

    // Read the response from server
    responseBody, err := io.ReadAll(response.Body)
    if err != nil {
        return nil, err
    }

    fmt.Println("Submitted ChainCode is: ", ChainCode)
    return responseBody, nil

}