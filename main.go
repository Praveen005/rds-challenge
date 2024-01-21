package main

import "fmt"

func main() {
	getDataFunc()
	resp, err := SubmitChaincode(submitURL, ChainCode)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(resp))
}
