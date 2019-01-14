package main

import (
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
		Owner:       "admin@vendor.pl",
	}
	authRequest.createSha()

	client := http.Client{Timeout: 30 * time.Second}

	tt := callHasContact(authRequest, client)
	log.Println(tt)

}
