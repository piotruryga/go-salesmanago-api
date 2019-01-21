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

	body, error := request.PrepareBody(authRequest, request.Email)
	req, error := request.PrepareRequest(body)
	t := time.Now()
	var timing int64 = 0
	result := true
	var reasonCode ReasonCode
	reasonCode = OK
	response, error := client.Do(req)
	if error != nil {
		log.Printf("Cannot post for method hasContact")
		result = false
		reasonCode = TIMEOUT
	}
	defer func() {
		if response != nil {
			response.Body.Close()
			result = true
		}
	}()
	if response != nil {
		timing = timeTrack(t, "hasContact")
		hasContactResponse, error := ParseResponse(response)
		if error != nil {
			log.Printf("Cannot parse response for method hasContact")
			result = false
			reasonCode = PARSE_ERROR
		}
		log.Printf("Has contact status: %v, message: %v, contactId: %v, reasonCode: %v", hasContactResponse.Success, hasContactResponse.Message,
			hasContactResponse.ContactId, reasonCode.String())
	}

	return New("hasContact", "1", timing, result, reasonCode.String())

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

func (request *HasContactRequest) PrepareRequest(body *bytes.Buffer) (*http.Request, error) {
	req, error := http.NewRequest("POST", APP1_ENDPOINT+"/contact/hasContact", body)
	req.Header.Set("Accept", HEADER_ACCEPT)
	req.Header.Set("Content-Type", HEADER_CONTENT_TYPE)
	return req, error
}

func (hasContactRequest *HasContactRequest) PrepareBody(authrequest AuthRequest, email string) (*bytes.Buffer, error) {
	hasContactRequest.AuthRequest = authrequest
	hasContactRequest.InitHasContactRequest(email)
	body := new(bytes.Buffer)
	error := json.NewEncoder(body).Encode(hasContactRequest)
	if error != nil {
		log.Print(error)
	}
	return body, error
}
