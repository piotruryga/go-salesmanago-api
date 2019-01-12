package main

import (
	"crypto/sha1"
	"encoding/hex"
)

type Authrequest struct {
	ApiKey      string `json:"apiKey"`
	ClientId    string `json:"clientId"`
	Sha         string `json:"sha"`
	RequestTime string `json:"requestTime"`
	ApiSecret   string
}

func (a *Authrequest) AddAuth(apiKey string, clientId string, requestTime string, apiSecret string) {
	a.ApiKey = apiKey
	a.ClientId = clientId
	a.RequestTime = requestTime
	a.ApiSecret = apiSecret
	a.Sha = createSha(apiKey, clientId, apiSecret)
}

func createSha(apiKey string, clientId string, apiSecret string) string {
	hash := sha1.New()
	hash.Write([]byte(apiKey + clientId + apiSecret))
	sha1_hash := hex.EncodeToString(hash.Sum(nil))
	return string(sha1_hash)
}
