package main

import (
	"bytes"
	"net/http"
)

type ApiMethod interface {
	CallMethod(authRequest AuthRequest, client http.Client) TimeTrack

	PrepareHasContactRequest(body *bytes.Buffer) (*http.Request, error)
}
