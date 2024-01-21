package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func handleGetreq(w http.ResponseWriter, r* http.Request){
	type dataToSend struct{
		NumParts	int			`json:"numParts"`
		Message		string		`json:"message"`
	}
	data := dataToSend{
		NumParts: 21,
		Message: "Hello our top programmer Praveen",
	}

	res, err := json.Marshal(data)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	
}


func startTestServer()(*httptest.Server){
	ts := httptest.NewServer(http.HandlerFunc(handleGetreq))
	return ts
}

func TestGetData(t *testing.T){
	ts := startTestServer()
	defer ts.Close()
	data := getDataCount(ts.URL);
	if data != 21{
		t.Errorf("Expected: 21, Got: %d", data)
	}
}