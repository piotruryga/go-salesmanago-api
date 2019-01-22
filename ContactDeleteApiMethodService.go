package main

import (
	"bytes"
	"errors"
	"log"
	"net/http"
)

type ContactDeleteRequest struct {
	AuthRequest
	Email string `json:"email"`
}

func (c *ContactDeleteRequest) InitContactDeleteRequest(email string) {
	c.Email = email
}

func (request *ContactDeleteRequest) CallMethod(authRequest AuthRequest, client http.Client) TimeTrack {
	//todo implement
	return TimeTrack{}
}

func (request *ContactDeleteRequest) PrepareBody(authrequest AuthRequest, email string) (*bytes.Buffer, error) {
	//todo implement
	return nil, errors.New("todo")
}

func (request *ContactDeleteRequest) PrepareRequest(body *bytes.Buffer) (*http.Request, error) {
	//todo implement
	return nil, nil
}

func callDeleteContact(authRequest AuthRequest, client http.Client) (TimeTrack, error) {
	if request, ok := ReturnImplementation("contactDeleteRequest").(*ContactDeleteRequest); ok {
		request.InitContactDeleteRequest("piotrek.uryga@gmail.com")
		return request.CallMethod(authRequest, client), nil
	} else {
		return TimeTrack{}, errors.New("cannot call hasContactRequest")
	}
}

func handleDeleteContact(authRequest AuthRequest, client http.Client) {
	deleteContactTt, error := callDeleteContact(authRequest, client)
	if error != nil {
		log.Printf("Cannot call deleteContact %s", error)
	}
	log.Println(deleteContactTt)
}
