package main

import (
	"bytes"
	"errors"
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
