// package main

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// )

// func fetchRemoteResource(url string) ([]byte , error){
// 	r, err := http.Get(url)

// 	if err != nil{
// 		return nil, err
// 	}

// 	defer r.Body.Close()
// 	return io.ReadAll(r.Body)
// }

// func main(){
// 	if len(os.Args) != 2{
// 		fmt.Fprintf(os.Stdout, "Must Specify a HTTP URL to get data from\n")
// 		os.Exit(1)
// 	}
// 	body, err := fetchRemoteResource(os.Args[1])

// 	if err != nil{
// 		fmt.Fprintf(os.Stdout, "%v\n", err)
// 		os.Exit(1)
// 	}
// 	fmt.Fprintf(os.Stdout, "%v\n", body)
// 	// os.Stdout.Write(body)
// }



/*

package main

import (
	"fmt"
	"io"
	"net/http"
	// "strings"
	"time"
)

const baseURL = "https://exam.ankush.wiki/data?part="
// const submitURL = "https://exam.ankush.wiki/answers"

func downloadData(partNumber int) ([]byte, error) {
	url := fmt.Sprintf("%s%d", baseURL, partNumber)
	//setting the header
	cookieHeader := "connect.sid=s%3AShnZDFr0rDdh-MPuhuym5ZwloAdckieu.PCQvfvexZw5CISOiluJdiPc%2BZNQ8UpQo%2F%2BEn%2Fe%2B7eI0"

	response, err := http.Get(url)
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
/*
func submitChaincode(chaincode string) error {
	payload := fmt.Sprintf(`{ "chaincode": "%s" }`, chaincode)
	response, err := http.Post(submitURL, "application/json", strings.NewReader(payload))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Process the response as needed

	return nil
}
*/
/*
func main() {
	// Iterate over all parts and download data
	for partNumber := 1; partNumber <= 23; partNumber++ {
		data, err := downloadData(partNumber)
		if err != nil {
			fmt.Println("Error downloading part", partNumber, ":", err)
			return
		}

		// Process the downloaded data (you might want to save it or analyze it)
		fmt.Println("Downloaded data for part", partNumber, ":", string(data))

		// Sleep to adhere to the API constraints (2.5 to 10 seconds apart)
		time.Sleep(time.Second * time.Duration(3)) // Adjust sleep duration as needed
	}

	// Create chaincode (implement your logic here)

	// Submit chaincode to /answers
	// replace "<YOUR_CHAINCODE>" with the actual chaincode value
	// err := submitChaincode("<YOUR_CHAINCODE>")
	// if err != nil {
	// 	fmt.Println("Error submitting chaincode:", err)
	// 	return
	// }

	fmt.Println("Mission completed successfully!")
}
*/

/*
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	// "strings"
	"time"
)

const (
	baseURL   = "https://exam.ankush.wiki/data?part="
	cookieHdr = "connect.sid=s%3AShnZDFr0rDdh-MPuhuym5ZwloAdckieu.PCQvfvexZw5CISOiluJdiPc%2BZNQ8UpQo%2F%2BEn%2Fe%2B7eI0"
	submitURL = "https://exam.ankush.wiki/answers"
)


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

/*
func SubmitChaincode(chaincode string) error {
	payload := fmt.Sprintf(`{ "chaincode": "%s" }`, chaincode)
	response, err := http.Post(submitURL, "text", strings.NewReader(payload))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Process the response as needed
	fmt.Println(response.Body)

	return nil
}
*/
/*
func main() {
	client := &http.Client{
		Timeout: time.Second * 10, // Example: Set timeout to 10 seconds
	}

	// Iterate over all parts and download data
	for partNumber := 1; partNumber <= 22; partNumber++ {
		data, err := downloadData(partNumber, client)
		if err != nil {
			log.Printf("Error downloading part %d: %v", partNumber, err)
			continue // Continue to the next part on error
		}

		// Process the downloaded data
		fmt.Printf("Downloaded data for part %d: %s\n", partNumber, string(data))
		JsonArray = append(JsonArray, string(data))

		// Sleep to adhere to the API constraints (2.5 to 10 seconds apart)
		time.Sleep(time.Second * 3) // Adjust sleep duration as needed
	}

	// Create chaincode (implement your logic here)


	/*
	// Submit chaincode to /answers
	// replace "<YOUR_CHAINCODE>" with the actual chaincode value
	var chainCode string
	fmt.Printf("\nPlease Enter the Chain code:\n")
	_, err := fmt.Scanf("%s", &chainCode)
	if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }
	err = SubmitChaincode(chainCode)
	if err != nil {
		fmt.Println("Error submitting chaincode:", err)
		return
	}
	*/
	/*
	ProcessJSONArray(JsonArray)
	fmt.Println("Mission completed successfully!")
}

*/



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
	submitURL = "https://exam.ankush.wiki/answers"
)


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

func main() {
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

	ProcessJSONArray(JsonArray)
	fmt.Println("Mission completed successfully!")
}
