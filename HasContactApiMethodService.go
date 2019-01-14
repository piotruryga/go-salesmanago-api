package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func callHasContact(authRequest Authrequest, client http.Client) TimeTrack {
	var hasContactRequest hasContactRequest
	body, error := prepareHasContactBody(&hasContactRequest, authRequest, "piotrek.uryga+7574@onet.pl")
	req, error := prepareHasContactRequest(body)
	t := time.Now()
	var timing int64 = 0
	response, error := client.Do(req)
	if error != nil {
		log.Fatalf("Cannot post")
	}
	defer func() {
		timing = timeTrack(t, "hasContact method")
		response.Body.Close()
	}()
	hasContactResponse, error := parseResponse(response)
	log.Println(hasContactResponse.Success, "", hasContactResponse.ContactId)
	return TimeTrack{
		MethodName:  "hasContact",
		AppInstance: "1",
		Time:        timing,
		Date:        time.Now(),
	}
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

func prepareHasContactBody(hasContactRequest *hasContactRequest, authrequest Authrequest, email string) (*bytes.Buffer, error) {
	hasContactRequest.Authrequest = authrequest
	hasContactRequest.InitHasContactRequest(email)
	body := new(bytes.Buffer)
	error := json.NewEncoder(body).Encode(hasContactRequest)
	if error != nil {
		log.Fatal(error)
	}
	return body, error
}
