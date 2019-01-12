package main

type authrequest struct {
	apiKey      string `json:"apiKey"`
	clientId    string `json:"clientId"`
	sha         string `json:"sha"`
	requestTime int64  `json:"requestTime"`
}
