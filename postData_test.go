package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func handlePostReq(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		defer r.Body.Close()
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			http.Error(w, "Invalid Content-Type", http.StatusBadRequest)
			return
		}
		data, err := io.ReadAll(r.Body)
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		fmt.Fprint(w, string(data))
		
	}else{
		http.Error(w, "Invalid HTTP method specified", http.StatusMethodNotAllowed)
	}
}

func startTestServerv2()*httptest.Server{
	ts := httptest.NewServer(http.HandlerFunc(handlePostReq))
	return ts
}

func TestSubmitChaincode(t *testing.T){
	ts := startTestServerv2()
	defer ts.Close()
	data := "praveen"
	resp, err := SubmitChaincode(ts.URL, data)
	if err != nil{
		t.Fatal(err)
	}
	if string(resp) != `{"chaincode":"praveen"}`{
		t.Errorf("Expected: %s, Got: %s", data, string(resp))
	}
}