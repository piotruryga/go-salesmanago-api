package main

import (
	"errors"
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

	var authRequest = AuthRequest{
		ApiKey:      "123abc",
		ClientId:    "h4jsu6pc5txybj04",
		RequestTime: strconv.FormatInt(time.Now().Unix(), 10),
		ApiSecret:   "4f69782e826841f794080cae87648e42",
		Owner:       "admin@vendor.pl",
	}
	authRequest.CreateSha()
	client := http.Client{Timeout: 30 * time.Second}
	InitRequestFactory()

	hasContactTt, error := callHasContact(authRequest, client)
	if error != nil {
		log.Fatalf("Cannot call hasContact")
	}
	log.Println(hasContactTt)

	deleteContactTt, error := callDeleteContact(authRequest, client)
	if error != nil {
		log.Fatalf("Cannot call deleteContact")
	}
	log.Println(deleteContactTt)

}

func callDeleteContact(authRequest AuthRequest, client http.Client) (TimeTrack, error) {
	if request, ok := ReturnImplementation("contactDeleteRequest").(*ContactDeleteRequest); ok {
		request.InitContactDeleteRequest("piotrek.uryga@gmail.com")
		return request.CallMethod(authRequest, client), nil
	} else {
		return TimeTrack{}, errors.New("cannot call hasContactRequest")
	}
}

func callHasContact(authRequest AuthRequest, client http.Client) (TimeTrack, error) {
	if request, ok := ReturnImplementation("hasContactRequest").(*HasContactRequest); ok {
		request.InitHasContactRequest("piotrek.uryga@gmail.com")
		return request.CallMethod(authRequest, client), nil
	} else {
		return TimeTrack{}, errors.New("cannot call hasContactRequest")
	}
}
