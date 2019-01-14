package main

import (
	"fmt"
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
	authRequest.createSha()

	client := http.Client{Timeout: 30 * time.Second}

	//var hasContactRequest HasContactRequest
	//tt := hasContactRequest.CallMethod(authRequest, client)
	//log.Println(tt)

	InitRF()

	if request, ok := ReturnImplementation("hasContactRequest").(*HasContactRequest); ok {
		//fmt.Println(ok)
		request.CallMethod(authRequest, client)
	} else {
		fmt.Println("XXX")
	}

}
