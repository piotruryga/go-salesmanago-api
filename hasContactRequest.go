package main

type hasContactRequest struct {
	Authrequest
	Email string `json:"email"`
}

func (h *hasContactRequest) InitHasContactRequest(email string) {
	h.Email = email
}
