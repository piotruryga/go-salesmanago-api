package main

type ContactDeleteRequest struct {
	Email string `json:"email"`
}

func (c *ContactDeleteRequest) InitContactDeleteRequest(email string) {
	c.Email = email
}
