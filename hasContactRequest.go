package main

type hasContactRequest struct {
	Authrequest
	Owner string `json:"owner"`
	Email string `json:"email"`
}

func (h *hasContactRequest) InitHasContactRequest(owner string, email string) {
	h.Owner = owner
	h.Email = email
}
