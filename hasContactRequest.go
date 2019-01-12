package main

import (
	"crypto/sha1"
	"encoding/hex"
)

type hasContactRequest struct {
	ApiKey      string `json:"apiKey"`
	ClientId    string `json:"clientId"`
	Sha         string `json:"sha"`
	RequestTime string `json:"requestTime"`
	ApiSecret   string

	Owner string `json:"owner"`
	Email string `json:"email"`
}

func (h *hasContactRequest) InitHasContactRequest(owner string, email string) {
	h.Owner = owner
	h.Email = email
}

func (h *hasContactRequest) AddAuth(apiKey string, clientId string, requestTime string, apiSecret string) {

	//should be interface
	h.ApiKey = apiKey
	h.ClientId = clientId
	h.RequestTime = requestTime
	h.ApiSecret = apiSecret
	h.Sha = createSha(apiKey, clientId, apiSecret)
}

func createSha(apiKey string, clientId string, apiSecret string) string {
	hash := sha1.New()
	hash.Write([]byte(apiKey + clientId + apiSecret))
	sha1_hash := hex.EncodeToString(hash.Sum(nil))
	return string(sha1_hash)
}
