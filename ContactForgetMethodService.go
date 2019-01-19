package main

import (
	"bytes"
	"net/http"
)

type ContactForgetRequest struct {
	AuthRequest
	Email string `json:"email"`
}

func (c *ContactForgetRequest) InitForgetRequest(email string) {
	c.Email = email
}

func (c *ContactForgetRequest) CallMethod(authRequest AuthRequest, client http.Client) TimeTrack {
	//todo implement
	return TimeTrack{}
}

func (c *ContactForgetRequest) PrepareBody(authRequest AuthRequest, email string) (*bytes.Buffer, error) {
	//todo implement
	return nil, nil
}

func (c *ContactForgetRequest) PrepareRequest(body *bytes.Buffer) (*http.Request, error) {
	//todo implement
	return nil, nil
}
