package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	APP1_ENDPOINT       = "https://www.salesmanago.pl/api"
	HEADER_ACCEPT       = "application/json, application/json"
	HEADER_CONTENT_TYPE = "application/json;charset=UTF-8"
)

func main() {

	var hasContactRequest hasContactRequest
	hasContactRequest.AddAuth("123abc", "h4jsu6pc5txybj04", strconv.FormatInt(time.Now().Unix(), 10), "4f69782e826841f794080cae87648e42")
	hasContactRequest.InitHasContactRequest("admin@vendor.pl", "piotrek.uryga+7574@onet.pl")

	body := new(bytes.Buffer)

	error := json.NewEncoder(body).Encode(hasContactRequest)
	if error != nil {
		log.Fatal(error)
	}

	req, error := http.NewRequest("POST", APP1_ENDPOINT+"/contact/hasContact", body)
	req.Header.Set("Accept", HEADER_ACCEPT)
	req.Header.Set("Content-Type", HEADER_CONTENT_TYPE)
	client := http.Client{Timeout: 30 * time.Second}
	response, error := client.Do(req)
	if error != nil {
		log.Fatalf("Cannot post")
	}
	defer response.Body.Close()
	bodyBytes, error := ioutil.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	var s = new(hasContactResponse)
	err := json.Unmarshal(bodyBytes, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	fmt.Println(bodyString)

}
