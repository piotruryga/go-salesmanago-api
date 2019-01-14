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

	var authRequest = Authrequest{
		ApiKey:      "123abc",
		ClientId:    "h4jsu6pc5txybj04",
		RequestTime: strconv.FormatInt(time.Now().Unix(), 10),
		ApiSecret:   "4f69782e826841f794080cae87648e42",
	}
	authRequest.createSha()

	client := http.Client{Timeout: 30 * time.Second}

	var hasContactRequest hasContactRequest
	body, error := prepareHasContactBody(&hasContactRequest, authRequest, "admin@vendor.pl", "piotrek.uryga+7574@onet.pl")
	req, error := prepareHasContactRequest(body)

	t := time.Now()
	response, error := client.Do(req)
	if error != nil {
		log.Fatalf("Cannot post")
	}
	defer func() {
		timeTrack(t, "hasContact method")
		response.Body.Close()
	}()

	hasContactResponse, error := parseResponse(response)
	log.Println(hasContactResponse.Success, "", hasContactResponse.ContactId)

}

func parseResponse(response *http.Response) (*HasContactResponse, error) {
	bodyBytes, error := ioutil.ReadAll(response.Body)
	var s = new(HasContactResponse)
	err := json.Unmarshal(bodyBytes, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}

	return s, error
}

func prepareHasContactRequest(body *bytes.Buffer) (*http.Request, error) {
	req, error := http.NewRequest("POST", APP1_ENDPOINT+"/contact/hasContact", body)
	req.Header.Set("Accept", HEADER_ACCEPT)
	req.Header.Set("Content-Type", HEADER_CONTENT_TYPE)
	return req, error
}

func prepareHasContactBody(hasContactRequest *hasContactRequest, authrequest Authrequest, owner string, email string) (*bytes.Buffer, error) {
	hasContactRequest.Authrequest = authrequest
	hasContactRequest.InitHasContactRequest(owner, email)
	body := new(bytes.Buffer)
	error := json.NewEncoder(body).Encode(hasContactRequest)
	if error != nil {
		log.Fatal(error)
	}
	return body, error
}
