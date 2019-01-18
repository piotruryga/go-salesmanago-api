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

type HasContactRequest struct {
	AuthRequest
	Email string `json:"email"`
}

func (h *HasContactRequest) InitHasContactRequest(email string) {
	h.Email = email
}

func (request *HasContactRequest) CallMethod(authRequest AuthRequest, client http.Client) TimeTrack {

	body, error := prepareHasContactBody(request, authRequest, request.Email)
	req, error := request.PrepareHasContactRequest(body)
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
	hasContactResponse, error := ParseResponse(response)
	log.Printf("Has contact status: %v, message: %v, contactId: %v", hasContactResponse.Success, hasContactResponse.Message, hasContactResponse.ContactId)
	return New("hasContact", "1", timing)

}

func ParseResponse(response *http.Response) (*HasContactResponse, error) {
	bodyBytes, error := ioutil.ReadAll(response.Body)
	var s = new(HasContactResponse)
	err := json.Unmarshal(bodyBytes, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}

	return s, error
}

func (request *HasContactRequest) PrepareHasContactRequest(body *bytes.Buffer) (*http.Request, error) {
	req, error := http.NewRequest("POST", APP1_ENDPOINT+"/contact/hasContact", body)
	req.Header.Set("Accept", HEADER_ACCEPT)
	req.Header.Set("Content-Type", HEADER_CONTENT_TYPE)
	return req, error
}

func prepareHasContactBody(hasContactRequest *HasContactRequest, authrequest AuthRequest, email string) (*bytes.Buffer, error) {
	hasContactRequest.AuthRequest = authrequest
	hasContactRequest.InitHasContactRequest(email)
	body := new(bytes.Buffer)
	error := json.NewEncoder(body).Encode(hasContactRequest)
	if error != nil {
		log.Fatal(error)
	}
	return body, error
}
