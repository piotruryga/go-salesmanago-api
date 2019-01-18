package main

import "net/http"

type RequestFactory interface {
	CallMethod(authRequest AuthRequest, client http.Client) TimeTrack
}

type ContactDeleteRequest struct {
	AuthRequest
	Email string `json:"email"`
}

var database map[string]interface{}

func InitRequestFactory() {
	database = make(map[string]interface{})
	hcR := new(HasContactRequest)
	database["hasContactRequest"] = hcR

	cdR := new(ContactDeleteRequest)
	database["contactDeleteRequest"] = cdR
}

func ReturnImplementation(requestName string) interface{} {
	return database[requestName]
}
